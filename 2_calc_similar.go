package main

import (
	"encoding/json"
	"fmt"
	"github.com/caneroj1/stemmer"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"github.com/james-bowman/sparse"
	"github.com/similar-manga/similar/mangadex"
	"github.com/similar-manga/similar/similar"
	"gonum.org/v1/gonum/mat"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {

	// Directory configuration
	// 39147 mangas in total as of 03-19-2022
	// 68308 mangas in total as of 04-03-2023
	dirData := "D:/MANGADEX/similar_data/"
	algoNumMin := -1
	algoNumMax := -1
	if len(os.Args) == 2 {
		dirData = os.Args[1]
	}
	if len(os.Args) == 4 {
		dirData = os.Args[1]
		algoNumMin, _ = strconv.Atoi(os.Args[2])
		algoNumMax, _ = strconv.Atoi(os.Args[3])
	}
	fmt.Printf("directory %s\n", dirData)
	dirMangas := dirData + "manga/"
	dirSimilar := dirData + "similar/"

	// Check that we have valid range
	if algoNumMin != -1 && algoNumMin >= algoNumMax {
		log.Fatalf("invalid range of %d to %d\n", algoNumMin, algoNumMax)
	}

	// Settings
	numSimToGet := 17
	tagScoreRatio := 0.40
	ignoreDescScoreUnder := 0.01
	acceptDescScoreOver := 0.45
	ignoreTagsUnderCount := 2
	minDescriptionWords := 15
	err := os.MkdirAll(dirSimilar, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Loop through all manga and try to get their chapter information for each
	countMangasProcessed := 0
	startProcessing := time.Now()
	corpusTag := []string{}
	corpusDesc := []string{}
	corpusDescLength := []int{}
	mangas := []mangadex.Manga{}
	itemsManga, _ := ioutil.ReadDir(dirMangas)
	for ct, file := range itemsManga {

		// If we are only updating a range, then skip mangas
		if algoNumMin != -1 && (ct < algoNumMin || ct >= algoNumMax) {
			continue
		}

		// Skip if a directory
		if file.IsDir() {
			continue
		}

		// Load the json from file into our manga struct
		manga := mangadex.Manga{}
		fileManga, _ := ioutil.ReadFile(dirMangas + file.Name())
		err := json.Unmarshal(fileManga, &manga)
		if err != nil {
			fmt.Printf("MANGA LOAD ERROR: %v (file %s)\n", err, file.Name())
			continue
		}

		// Skip if invalid
		if manga.Attributes.Title == nil || manga.Attributes.Description == nil {
			continue
		}

		// Get the tag and description for this manga
		tagText := ""
		for _, tag := range manga.Attributes.Tags {
			if tag.Type_ != "tag" {
				continue
			}
			reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
			tagText += reg.ReplaceAllString((*tag.Attributes.Name)["en"], "") + " "
		}
		descText := similar.CleanTitle((*manga.Attributes.Title)["en"]) + " "
		for _, altTitle := range manga.Attributes.AltTitles {
			if val, ok := altTitle["en"]; ok {
				if similar.CleanTitle(val) != "" {
					descText += similar.CleanTitle(val) + " "
				}
			}
		}
		descText += similar.CleanDescription((*manga.Attributes.Description)["en"])

		// Append to the corpusDesc
		corpusTag = append(corpusTag, tagText)
		corpusDesc = append(corpusDesc, descText)
		corpusDescLength = append(corpusDescLength, len(strings.Split(descText, " ")))
		mangas = append(mangas, manga)

		// Debug
		if len(mangas)%1000 == 0 {
			fmt.Printf("%d/%d mangas loaded....\n", len(mangas), len(itemsManga))
		}
		//fmt.Printf("%s - volumn %s - chapter %s\n", manga.Data.Id, (*manga.Data.Attributes).LastVolume, (*manga.Data.Attributes).LastChapter)
		//if len(mangas) >= 5000 {
		//	break
		//}

	}
	fmt.Printf("loaded %d magas in our corupus\n", len(corpusDesc))

	// Create our tf-idf pipeline
	lsiTagVectoriser := nlp.NewCountVectoriser([]string{}...)
	lsiPipelineTag := nlp.NewPipeline(lsiTagVectoriser)
	stopWordsStemmed := append([]string(nil), similar.StopWords...)
	stemmer.StemMultipleMutate(&stopWordsStemmed)
	for i := range stopWordsStemmed {
		stopWordsStemmed[i] = strings.ToLower(stopWordsStemmed[i])
	}
	lsiPipelineDescription := nlp.NewPipeline(nlp.NewCountVectoriser(stopWordsStemmed...), nlp.NewTfidfTransformer())

	// Transform the corpusTag into an LSI fitting the model to the documents in the process
	start := time.Now()
	fmt.Printf("fitting to corpus of tags!\n")
	lsiTag, err := lsiPipelineTag.FitTransform(corpusTag...)
	if err != nil {
		log.Fatalf("ERROR: failed to process documents because\n %v\n", err)
	}
	lsiTagCSC := lsiTag.(sparse.TypeConverter).ToCSC()
	m, n := lsiTag.Dims()
	fmt.Printf("\t- fitted data in %s\n", time.Since(start))
	fmt.Printf("\t- system dim = %d x %d\n\n", m, n)

	// We will now apply our custom weights for tags
	// Each row of this matrix is a tag which we have a weight for
	fmt.Println("Tag Vectoriser Vocabulary:")
	fmt.Println(lsiTagVectoriser.Vocabulary)
	fmt.Println()
	vocabularyInverse := map[int]string{}
	for k, v := range lsiTagVectoriser.Vocabulary {
		vocabularyInverse[v] = k
	}

	// Special weights for tags that should have higher priority over others
	// These are hand tuned and adhoc in nature, but seem to work?
	tagWeights := map[string]float64{
		"sexualviolence": 1.00,
		"gore":           1.00,
		"koma":           1.00,
		"wuxia":          1.00,
		"loli":           0.90,
		"incest":         0.90,
		"sports":         0.90,
		"boyslove":       0.90,
		"girlslove":      0.90,
		"isekai":         0.90,
		"villainess":     0.90,
		"historical":     0.80,
		"horror":         0.80,
		"mecha":          0.80,
		"medical":        0.80,
		"sliceoflife":    0.80,
		"cooking":        0.80,
		"crossdressing":  0.80,
		"genderswap":     0.80,
		"harem":          0.80,
		"reverseharem":   0.80,
		"vampires":       0.80,
		"zombies":        0.80,
	}

	// Loop through the tag weights and set them to our custom ones
	lsiTagCSCWeighted := lsiTag.(sparse.TypeConverter).ToCSC()
	dimR, dimC := lsiTagCSCWeighted.Dims()
	for r := 0; r < dimR; r++ {
		tag := vocabularyInverse[r]
		tagWeight := 0.70
		if val, ok := tagWeights[tag]; ok {
			tagWeight = val
		}
		for c := 0; c < dimC; c++ {
			if lsiTagCSCWeighted.At(r, c) > 0 {
				lsiTagCSCWeighted.Set(r, c, tagWeight)
			}
		}
	}

	// Transform the corpusDesc into an LSI fitting the model to the documents in the process
	start = time.Now()
	fmt.Printf("fitting to corpus of descriptions!\n")
	lsiDesc, err := lsiPipelineDescription.FitTransform(corpusDesc...)
	if err != nil {
		log.Fatalf("ERROR: failed to process documents because\n %v\n", err)
	}
	lsiDescCSC := lsiDesc.(sparse.TypeConverter).ToCSC()
	m, n = lsiDesc.Dims()
	fmt.Printf("\t- fitted data in %s\n", time.Since(start))
	fmt.Printf("\t- system dim = %d x %d\n\n", m, n)

	// Create a "buffer" that is our num of max rutines
	// If we can append to it, then we will run a coroutine
	// https://stackoverflow.com/a/25306241/7718197
	// https://downey.io/notes/dev/openmp-parallel-for-in-golang/
	var wg sync.WaitGroup
	wg.Add(len(mangas))
	maxGoroutines := 6
	guard := make(chan struct{}, maxGoroutines)
	var mu sync.Mutex

	// For each manga we will get the top similar for tags and description
	// We will then combine these into a single score which is then used to rank all manga
	// TODO: skip matched manga that are already a "related" manga list
	start = time.Now()
	for j := 0; j < len(mangas); j++ {

		// would block if guard channel is already filled
		guard <- struct{}{}
		go func(j int) {
			defer wg.Done()

			// This manga we will try to match to
			// NOTE: here we use the weighted tag CSC matrix, so we will multiply this against a one-hot-matrix
			// NOTE: e.g. [0.7 1.0 0.0 0.0 0.9] * [0 1 0 0 1] => 1.9 score value for current against another
			manga := mangas[j]
			vTagWeighted := lsiTagCSCWeighted.ColView(j)
			numTags := int(mat.Sum(lsiTagCSC.ColView(j)))
			vDesc := lsiDescCSC.ColView(j)

			// Skip this manga if it has no description
			if corpusDescLength[j] < minDescriptionWords {
				<-guard
				return
			}

			// Debug check / skip mangas
			//debugMangaIds := map[string]bool{"e56a163f-1a4c-400b-8c1d-6cb98e63ce04": true}
			//debugMangaIds := map[string]bool{"ee0df4ab-1e8d-49b9-9404-da9dcb11a32a": true}
			//debugMangaIds := map[string]bool{"32d76d19-8a05-4db0-9fc2-e0b0648fe9d0": true, "d46d9573-2ad9-45b2-9b6d-45f95452d1c0": true,
			//	"e78a489b-6632-4d61-b00b-5206f5b8b22b": true, "58bc83a0-1808-484e-88b9-17e167469e23": true, "0fa5dab2-250a-4f69-bd15-9ceea54176fa": true}
			//if _, ok := debugMangaIds[manga.Id]; !ok {
			//	<-guard
			//	return
			//}

			// Type of match which also stores the description
			// Modeled after nlp.Match object
			type CustomMatch struct {
				ID           interface{}
				Distance     float64
				DistanceTag  float64
				DistanceDesc float64
			}

			// Perform matching to all the other vectors
			var matches []CustomMatch
			for k := 0; k < len(mangas); k++ {

				// Get score for both tags and description
				distTag := pairwise.CosineSimilarity(vTagWeighted, lsiTagCSC.ColView(k))
				distDesc := pairwise.CosineSimilarity(vDesc, lsiDescCSC.ColView(k))

				// Reject invalid matches
				if math.IsNaN(distTag) || distTag < 1e-4 {
					distTag = 0
				}
				if math.IsNaN(distDesc) || distDesc < 1e-4 {
					distDesc = 0
				}

				// Special reject criteria to try to be robust to small label / description length
				if numTags < ignoreTagsUnderCount {
					distTag = 1
				}
				if distDesc < ignoreDescScoreUnder || corpusDescLength[k] < minDescriptionWords {
					distDesc = 0
				}
				if distDesc > acceptDescScoreOver {
					distTag = 1
				}

				// Combine the two
				match := CustomMatch{}
				match.ID = k
				match.Distance = tagScoreRatio*distTag + distDesc
				match.DistanceTag = distTag
				match.DistanceDesc = distDesc
				matches = append(matches, match)

			}
			sort.Slice(matches, func(i, j int) bool {
				return matches[i].Distance > matches[j].Distance
			})

			// Create our similar manga api object which will have our matches in it
			similarMangaData := similar.SimilarManga{}
			similarMangaData.Id = manga.Id
			similarMangaData.Title = *manga.Attributes.Title
			similarMangaData.ContentRating = manga.Attributes.ContentRating
			similarMangaData.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05+00:00")
			//fmt.Printf("manga %d has %d tags -> %s - https://mangadex.org/title/%s\n", j, numTags, (*manga.Attributes.Title)["en"], manga.Id)

			// Finally loop through all our matches and try to find the best ones!
			var matches_best []CustomMatch
			for _, match := range matches {

				// Skip if not a valid score
				if match.Distance <= 0 {
					continue
				}

				// Skip if the same id
				id := match.ID.(int)
				if id == j {
					continue
				}

				// Skip if no chapters
				if mangas[id].Attributes.LastChapter != "" {
					//fmt.Printf("\u001B[1;33m\t - match %d has no translated chapters! -> %s\u001B[0m\n", id, (*mangas[id].Attributes.Title)["en"])
					continue
				}

				// Skip if no common languages
				// This also enforces that the other manga has at least one chapter a user can read!
				foundCommonLang := false
				for _, lang1 := range manga.Attributes.AvailableTranslatedLanguages {
					for _, lang2 := range mangas[id].Attributes.AvailableTranslatedLanguages {
						if lang1 == lang2 {
							foundCommonLang = true
						}
						if foundCommonLang {
							break
						}
					}
					if foundCommonLang {
						break
					}
				}
				if !foundCommonLang && len(manga.Attributes.AvailableTranslatedLanguages) > 0 {
					//fmt.Printf("\u001B[1;33m\t - match %d had no commmon lang! -> %s (%s) https://mangadex.org/title/%s\u001B[0m\n",
					//	id, (*mangas[id].Attributes.Title)["en"], strings.Join(mangas[id].Attributes.AvailableTranslatedLanguages, ","), mangas[id].Id)
					continue
				}

				// Tags / content ratings / demographics we enforce
				// Also enforce that the manga can't be *related* to the match
				if similar.NotValidMatch(manga, mangas[id]) {
					continue
				}

				// Otherwise lets append it!
				matchData := similar.SimilarMatch{}
				matchData.Id = mangas[id].Id
				matchData.Title = *mangas[id].Attributes.Title
				matchData.ContentRating = mangas[id].Attributes.ContentRating
				matchData.Score = float32(match.Distance) / float32(tagScoreRatio+1.0)
				matchData.Languages = mangas[id].Attributes.AvailableTranslatedLanguages
				similarMangaData.SimilarMatches = append(similarMangaData.SimilarMatches, matchData)
				matches_best = append(matches_best, match)
				//fmt.Printf("\t - matched to id %d (%.3f tag, %.3f desc, %.3f combined) -> %s - https://mangadex.org/title/%s\n",
				//	id, match.DistanceTag, match.DistanceDesc, matchData.Score, (*mangas[id].Attributes.Title)["en"], mangas[id].Id)

				// Debug error if score is invalid
				if matchData.Score > 1 || matchData.Score < 0 {
					log.Fatalf("\u001B[1;31mINVALID SCORE: %s -> %s gave %.4f\u001B[0m\n", similarMangaData.Id, mangas[id].Id, matchData.Score)
				}

				// Exit if we have found enough similar manga!
				if len(similarMangaData.SimilarMatches) >= numSimToGet {
					break
				}

			}

			// Finally if we have non-zero matches then we should save it!
			if len(similarMangaData.SimilarMatches) > 0 {
				file, _ := json.MarshalIndent(similarMangaData, "", " ")
				_ = ioutil.WriteFile(dirSimilar+similarMangaData.Id+".json", file, 0644)
			}
			countMangasProcessed++
			avgIterTime := float64(j+1) / time.Since(start).Seconds()
			{
				mu.Lock()
				defer mu.Unlock()
				fmt.Printf("manga %d has %d tags -> %s - https://mangadex.org/title/%s\n", j, numTags, (*manga.Attributes.Title)["en"], manga.Id)
				for i, match := range matches_best {
					id := match.ID.(int)
					score := similarMangaData.SimilarMatches[i].Score
					fmt.Printf("  - matched %d (%.3f tag, %.3f desc, %.3f comb) -> %s - https://mangadex.org/title/%s\n",
						id, match.DistanceTag, match.DistanceDesc, score, (*mangas[id].Attributes.Title)["en"], mangas[id].Id)
				}
				fmt.Printf("%d/%d processed at %.2f manga/sec....\n\n", j+1, len(mangas), avgIterTime)
			}
			<-guard
		}(j)

	}
	wg.Wait()
	fmt.Printf("calculated simularity for %d mangas in %s!!\n\n", countMangasProcessed, time.Since(startProcessing))

}
