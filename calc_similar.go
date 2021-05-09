package main

import (
	"./mangadex"
	"./similar"
	"encoding/json"
	"fmt"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
	_ "gonum.org/v1/gonum/mat"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

func main() {

	// Directory configuration
	dirMangas := "../data/manga/"
	dirSimilar := "../data/similar/"
	numSimToGet := 5
	err := os.MkdirAll(dirSimilar, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Loop through all manga and try to get their chapter information for each
	corpus := []string{}
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
		_ = json.Unmarshal(fileManga, &manga)

		// Append to the corpus
		corpus = append(corpus, manga.Data.Attributes.Description["en"])
		mangas = append(mangas, manga)

		// Debug
		if len(mangas)%200 == 0 {
			fmt.Printf("%d/%d mangas loaded....\n", len(mangas), len(itemsManga))
		}
		if len(mangas) > 300 {
			break
		}

	}
	fmt.Printf("loaded %d magas in our corupus\n", len(corpus))

	// Create our tf-idf pipeline
	var stopWords = []string{"a", "about", "above", "above", "across", "after", "afterwards", "again", "against",
		"all", "almost", "alone", "along", "already", "also", "although", "always", "am", "among", "amongst", "amoungst", "amount", "an", "and", "another", "any", "anyhow", "anyone", "anything", "anyway", "anywhere", "are", "around", "as", "at", "back", "be", "became", "because", "become", "becomes", "becoming", "been", "before", "beforehand", "behind", "being", "below", "beside", "besides", "between", "beyond", "bill", "both", "bottom", "but", "by", "call", "can", "cannot", "cant", "co", "con", "could", "couldnt", "cry", "de", "describe", "detail", "do", "done", "down", "due", "during", "each", "eg", "eight", "either", "eleven", "else", "elsewhere", "empty", "enough", "etc", "even", "ever", "every", "everyone", "everything", "everywhere", "except", "few", "fifteen", "fify", "fill", "find", "fire", "first", "five", "for", "former", "formerly", "forty", "found", "four", "from", "front", "full", "further", "get", "give", "go", "had", "has", "hasnt", "have", "he", "hence", "her", "here", "hereafter", "hereby", "herein", "hereupon", "hers", "herself", "him", "himself", "his", "how", "however", "hundred", "ie", "if", "in", "inc", "indeed", "interest", "into", "is", "it", "its", "itself", "keep", "last", "latter", "latterly", "least", "less", "ltd", "made", "many", "may", "me", "meanwhile", "might", "mill", "mine", "more", "moreover", "most", "mostly", "move", "much", "must", "my", "myself", "name", "namely", "neither", "never", "nevertheless", "next", "nine", "no", "nobody", "none", "noone", "nor", "not", "nothing", "now", "nowhere", "of", "off", "often", "on", "once", "one", "only", "onto", "or", "other", "others", "otherwise", "our", "ours", "ourselves", "out", "over", "own", "part", "per", "perhaps", "please", "put", "rather", "re", "same", "see", "seem", "seemed", "seeming", "seems", "serious", "several", "she", "should", "show", "side", "since", "sincere", "six", "sixty", "so", "some", "somehow", "someone", "something", "sometime", "sometimes", "somewhere", "still", "such", "system", "take", "ten", "than", "that", "the", "their", "them", "themselves", "then", "thence", "there", "thereafter", "thereby", "therefore", "therein", "thereupon", "these", "they", "thickv", "thin", "third", "this", "those", "though", "three", "through", "throughout", "thru", "thus", "to", "together", "too", "top", "toward", "towards", "twelve", "twenty", "two", "un", "under", "until", "up", "upon", "us", "very", "via", "was", "we", "well", "were", "what", "whatever", "when", "whence", "whenever", "where", "whereafter", "whereas", "whereby", "wherein", "whereupon", "wherever", "whether", "which", "while", "whither", "who", "whoever", "whole", "whom", "whose", "why", "will", "with", "within", "without", "would", "yet", "you", "your", "yours", "yourself", "yourselves"}
	vectoriser := nlp.NewCountVectoriser(stopWords...)
	transformer := nlp.NewTfidfTransformer()
	lsiPipeline := nlp.NewPipeline(vectoriser, transformer)

	// Transform the corpus into an LSI fitting the model to the documents in the process
	start := time.Now()
	fmt.Printf("fitting to the corpus!\n")
	lsi, err := lsiPipeline.FitTransform(corpus...)
	if err != nil {
		fmt.Printf("failed to process documents because %v\n", err)
		return
	}
	fmt.Printf("fitted to corpus in %s\n", time.Since(start))
	m, n := lsi.Dims()
	fmt.Printf("lsi dim = %d x %d\n", m, n)

	// iterate over document feature vectors (columns) in the LSI matrix and compare
	// with the query vector for similarity. Similarity is determined by the difference
	// between the angles of the vectors known as the cosine similarity
	index := nlp.NewLinearScanIndex(pairwise.CosineDistance)
	nlp.ColDo(lsi, func(j int, v mat.Vector) {
		index.Index(v, j)
	})

	// For each manga get the top k similar
	nlp.ColDo(lsi, func(j int, v mat.Vector) {

		// Search to get the top k matches
		matches := index.Search(v, numSimToGet)
		sort.Slice(matches[:], func(i, j int) bool {
			return matches[i].Distance > matches[j].Distance
		})

		// Check to see if our score is nan (no valid match)
		hasNonNanScore := false
		for _, match := range matches {
			if !math.IsNaN(match.Distance) && match.Distance > 1e-4 {
				hasNonNanScore = true
			}
		}
		if !hasNonNanScore {
			return
		}

		// Finally print out our scores
		similarMangaData := similar.SimilarManga{}
		similarMangaData.Id = mangas[j].Data.Id
		similarMangaData.Title = mangas[j].Data.Attributes.Title
		similarMangaData.ContentRating = mangas[j].Data.Attributes.ContentRating
		similarMangaData.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05+00:00")
		fmt.Printf("manga %d received %d matches -> %s\n", j, len(matches), mangas[j].Data.Attributes.Title["en"])
		for _, match := range matches {

			// Skip if not a valid score
			if math.IsNaN(match.Distance) || match.Distance < 1e-4 {
				continue
			}

			// Otherwise lets append it!
			fmt.Printf("\t - matched to id %d (%.3f score) -> %s\n", match.ID.(int), match.Distance, mangas[match.ID.(int)].Data.Attributes.Title["en"])
			id := match.ID.(int)
			matchData := similar.SimilarMatch{}
			matchData.Id = mangas[id].Data.Id
			matchData.Title = mangas[id].Data.Attributes.Title
			matchData.ContentRating = mangas[id].Data.Attributes.ContentRating
			matchData.Score = float32(match.Distance)
			similarMangaData.SimilarMatches = append(similarMangaData.SimilarMatches, matchData)

		}

		// Finally if we have non-zero matches then we should save it!
		if len(similarMangaData.SimilarMatches) > 0 {
			file, _ := json.MarshalIndent(similarMangaData, "", " ")
			_ = ioutil.WriteFile(dirSimilar+similarMangaData.Id+".json", file, 0644)
		}

	})

}
