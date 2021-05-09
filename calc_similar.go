package main

import (
	"./mangadex"
	"./similar"
	"encoding/json"
	"fmt"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
	_ "gonum.org/v1/gonum/mat"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {

	// Directory configuration
	dirMangas := "../data/manga/"
	dirSimilar := "../data/similar/"
	numSimToGet := 10
	tagScoreRatio := 0.6
	err := os.MkdirAll(dirSimilar, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Loop through all manga and try to get their chapter information for each
	corpusTag := []string{}
	corpusDesc := []string{}
	mangas := []mangadex.MangaResponse{}
	itemsManga, _ := ioutil.ReadDir(dirMangas)
	for _, file := range itemsManga {

		// Skip if a directory
		if file.IsDir() {
			continue
		}

		// Load the json from file into our manga struct
		manga := mangadex.MangaResponse{}
		fileManga, _ := ioutil.ReadFile(dirMangas + file.Name())
		err := json.Unmarshal(fileManga, &manga)
		if err != nil {
			fmt.Printf("MANGA LOAD ERROR: %v (file %s)\n", err, file.Name())
			continue
		}

		// Get the tag and description for this manga
		tagText := ""
		for _, tag := range manga.Data.Attributes.Tags {
			if tag.Type_ != "tag" {
				continue
			}
			tagText += strings.ReplaceAll(tag.Attributes.Name["en"], " ", "") + " "
		}
		descText := manga.Data.Attributes.Description["title"] + " " + manga.Data.Attributes.Description["en"]

		// Append to the corpusDesc
		corpusTag = append(corpusTag, tagText)
		corpusDesc = append(corpusDesc, descText)
		mangas = append(mangas, manga)

		// Debug
		if len(mangas)%200 == 0 {
			fmt.Printf("%d/%d mangas loaded....\n", len(mangas), len(itemsManga))
		}
		//if len(mangas) >= 300 {
		//	break
		//}

	}
	fmt.Printf("loaded %d magas in our corupus\n", len(corpusDesc))

	// Create our tf-idf pipeline
	lsiPipelineTag := nlp.NewPipeline(nlp.NewCountVectoriser(), nlp.NewTfidfTransformer())
	lsiPipelineDescription := nlp.NewPipeline(nlp.NewCountVectoriser(similar.StopWords...), nlp.NewTfidfTransformer())

	// Transform the corpusTag into an LSI fitting the model to the documents in the process
	start := time.Now()
	fmt.Printf("fitting to corpus of tags!\n")
	lsiTag, err := lsiPipelineTag.FitTransform(corpusTag...)
	if err != nil {
		log.Fatalf("ERROR: failed to process documents because\n %v\n", err)
	}
	m, n := lsiTag.Dims()
	fmt.Printf("\t- fitted data in %s\n", time.Since(start))
	fmt.Printf("\t- system dim = %d x %d\n", m, n)

	// Transform the corpusDesc into an LSI fitting the model to the documents in the process
	start = time.Now()
	fmt.Printf("fitting to corpus of descriptions!\n")
	lsiDesc, err := lsiPipelineDescription.FitTransform(corpusDesc...)
	if err != nil {
		log.Fatalf("ERROR: failed to process documents because\n %v\n", err)
	}
	m, n = lsiDesc.Dims()
	fmt.Printf("\t- fitted data in %s\n", time.Since(start))
	fmt.Printf("\t- system dim = %d x %d\n", m, n)

	// Convert lsi matrices to dense so we can do column views...
	lsiTag = lsiTag.(sparse.TypeConverter).ToCSC()
	lsiDesc = lsiDesc.(sparse.TypeConverter).ToCSC()

	// For each manga we will get the top similar for tags and description
	// We will then combine these into a single score which is then used to rank all manga
	start = time.Now()
	for j := 0; j < len(mangas); j++ {

		// This manga we will try to match to
		manga :=  mangas[j]
		vTag := lsiTag.(mat.ColViewer).ColView(j)
		vDesc := lsiDesc.(mat.ColViewer).ColView(j)

		// Perform matching to all the other vectors
		var matches []nlp.Match
		for k := 0; k < len(mangas); k++ {

			// Get score for both tags and description
			distTag := pairwise.CosineSimilarity(vTag, lsiTag.(mat.ColViewer).ColView(k))
			distDesc := pairwise.CosineSimilarity(vDesc, lsiDesc.(mat.ColViewer).ColView(k))
			if math.IsNaN(distTag) || distTag < 1e-4 {
				distTag = 0
			}
			if math.IsNaN(distDesc) || distDesc < 1e-4 {
				distDesc = 0
			}

			// Combine the two
			match := nlp.Match{}
			match.ID = k
			match.Distance = tagScoreRatio*distTag + (1.0-tagScoreRatio)*distDesc
			matches = append(matches, match)

		}
		sort.Slice(matches[:], func(i, j int) bool {
			return matches[i].Distance > matches[j].Distance
		})

		// Create our similar manga api object which will have our matches in it
		similarMangaData := similar.SimilarManga{}
		similarMangaData.Id = manga.Data.Id
		similarMangaData.Title = manga.Data.Attributes.Title
		similarMangaData.ContentRating = manga.Data.Attributes.ContentRating
		similarMangaData.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05+00:00")
		fmt.Printf("manga %d -> %s\n", j, manga.Data.Attributes.Title["en"])
		fmt.Printf("%s\n",corpusTag[j])

		// Finally loop through all our matches and try to find the best ones!
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

			// Tags / content ratings / demographics we enforce
			if similar.NotValidMatch(manga, mangas[id]) {
				continue
			}

			// Otherwise lets append it!
			fmt.Printf("\t - matched to id %d (%.3f score) -> %s\n", id, match.Distance, mangas[id].Data.Attributes.Title["en"])
			fmt.Printf("\t - %s\n",corpusTag[id])
			matchData := similar.SimilarMatch{}
			matchData.Id = mangas[id].Data.Id
			matchData.Title = mangas[id].Data.Attributes.Title
			matchData.ContentRating = mangas[id].Data.Attributes.ContentRating
			matchData.Score = float32(match.Distance)
			similarMangaData.SimilarMatches = append(similarMangaData.SimilarMatches, matchData)

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
		avgIterTime := float64(j+1) / (1e-9 * float64(time.Since(start).Nanoseconds()))
		fmt.Printf("%d/%d processed at %.2f manga/sec....\n\n", j+1, len(mangas), avgIterTime)

	}


}
