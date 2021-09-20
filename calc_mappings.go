package main

import (
	"./external"
	"./mangadex"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func openCSVFileStream(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func writeToCSV(writer *csv.Writer, data []string) {
	err := writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// Directory configuration
	dirMangas := "../similar_data/manga/"
	dirMappings := "../similar_data/mapping/"
	updateAltCoverMapping := false
	err := os.MkdirAll(dirMappings, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// id, title, and content rating (for cache searching)
	fileSEARCH := openCSVFileStream(dirMappings + "mdex2search.csv")
	defer fileSEARCH.Close()
	writerSEARCH := csv.NewWriter(fileSEARCH)
	defer writerSEARCH.Flush()

	// anilist
	// https://anilist.co/manga/`{id}`
	fileAL := openCSVFileStream(dirMappings + "anilist2mdex.csv")
	defer fileAL.Close()
	writerAL := csv.NewWriter(fileAL)
	defer writerAL.Flush()

	// animeplanet
	// https://www.anime-planet.com/manga/`{slug}`
	fileAP := openCSVFileStream(dirMappings + "animeplanet2mdex.csv")
	defer fileAP.Close()
	writerAP := csv.NewWriter(fileAP)
	defer writerAP.Flush()

	// bookwalker.jp
	// https://bookwalker.jp/`{slug}`
	fileBW := openCSVFileStream(dirMappings + "bookwalker2mdex.csv")
	defer fileBW.Close()
	writerBW := csv.NewWriter(fileBW)
	defer writerBW.Flush()

	// mangaupdates
	// https://www.mangaupdates.com/series.html?id=`{id}`
	fileMU := openCSVFileStream(dirMappings + "mangaupdates2mdex.csv")
	defer fileMU.Close()
	writerMU := csv.NewWriter(fileMU)
	defer writerMU.Flush()

	// novelupdates
	// https://www.novelupdates.com/series/`{slug}`
	fileNU := openCSVFileStream(dirMappings + "novelupdates2mdex.csv")
	defer fileNU.Close()
	writerNU := csv.NewWriter(fileNU)
	defer writerNU.Flush()

	// kitsu.io
	// https://kitsu.io/api/edge/manga?filter[slug]={slug}
	fileKT := openCSVFileStream(dirMappings + "kitsu2mdex.csv")
	defer fileKT.Close()
	writerKT := csv.NewWriter(fileKT)
	defer writerKT.Flush()

	// myanimelist
	// https://myanimelist.net/manga/{id}
	fileMAL := openCSVFileStream(dirMappings + "myanimelist2mdex.csv")
	defer fileMAL.Close()
	writerMAL := csv.NewWriter(fileMAL)
	defer writerMAL.Flush()

	// Loop through all manga and try to get their chapter information for each
	start := time.Now()
	itemsManga, _ := ioutil.ReadDir(dirMangas)
	for ct, file := range itemsManga {

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

		// Our search file
		data := []string{manga.Id, (*manga.Attributes.Title)["en"], manga.Attributes.ContentRating}
		writeToCSV(writerSEARCH, data)

		// Save the external mappings
		if _, ok := manga.Attributes.Links["al"]; ok {
			data := []string{manga.Attributes.Links["al"], manga.Id}
			writeToCSV(writerAL, data)
		}
		if _, ok := manga.Attributes.Links["ap"]; ok {
			data := []string{manga.Attributes.Links["ap"], manga.Id}
			writeToCSV(writerAP, data)
		}
		if _, ok := manga.Attributes.Links["bw"]; ok {
			data := []string{manga.Attributes.Links["bw"], manga.Id}
			writeToCSV(writerBW, data)
		}
		if _, ok := manga.Attributes.Links["mu"]; ok {
			data := []string{manga.Attributes.Links["mu"], manga.Id}
			writeToCSV(writerMU, data)
		}
		if _, ok := manga.Attributes.Links["nu"]; ok {
			data := []string{manga.Attributes.Links["nu"], manga.Id}
			writeToCSV(writerNU, data)
		}
		if _, ok := manga.Attributes.Links["kt"]; ok {
			data := []string{manga.Attributes.Links["kt"], manga.Id}
			writeToCSV(writerKT, data)
		}
		if _, ok := manga.Attributes.Links["mal"]; ok {
			data := []string{manga.Attributes.Links["mal"], manga.Id}
			writeToCSV(writerMAL, data)
		}

		// Debug
		if ct%100 == 0 {
			avgIterTime := float64(ct+1) / time.Since(start).Seconds()
			fmt.Printf("%d/%d mangas -> processing at %.2f manga/sec....\n", ct+1, len(itemsManga), avgIterTime)
		}

	}
	fmt.Printf("done processing mappings!\n")
	writerSEARCH.Flush()
	writerAL.Flush()
	writerAP.Flush()
	writerBW.Flush()
	writerMU.Flush()
	writerNU.Flush()
	writerKT.Flush()
	writerMAL.Flush()
	fileSEARCH.Close()
	fileAL.Close()
	fileAP.Close()
	fileBW.Close()
	fileMU.Close()
	fileNU.Close()
	fileKT.Close()
	fileMAL.Close()

	// Alternative cover urls for use if we are cache'ing
	if updateAltCoverMapping {

		// Open save file
		fileAlternativeImage := openCSVFileStream(dirMappings + "mdex2altimage.csv")
		defer fileAlternativeImage.Close()
		writerAlternativeImage := csv.NewWriter(fileAlternativeImage)
		defer writerAlternativeImage.Flush()

		// Loop through all manga and try to get their chapter information for each
		start = time.Now()
		countHaveImagesExternal := make(map[string]int)
		countHaveImages := 0
		for ct, file := range itemsManga {

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

			// Get our url for this manga if we can
			url := ""
			if _, ok := manga.Attributes.Links["kt"]; ok {
				url = external.GetCoverKitsu(manga.Attributes.Links["kt"])
				countHaveImagesExternal["kt"]++
			}
			if _, ok := manga.Attributes.Links["al"]; url == "" && ok {
				url = external.GetCoverAniList(manga.Attributes.Links["al"])
				countHaveImagesExternal["al"]++
			}
			if _, ok := manga.Attributes.Links["mal"]; url == "" && ok {
				url = external.GetCoverMyAnimeList(manga.Attributes.Links["mal"])
				countHaveImagesExternal["mal"]++
			}
			if _, ok := manga.Attributes.Links["mu"]; url == "" && ok {
				url = external.GetCoverMangaUpdates(manga.Attributes.Links["mu"])
				countHaveImagesExternal["mu"]++
			}
			if _, ok := manga.Attributes.Links["ap"]; url == "" && ok {
				url = external.GetCoverAnimePlanet(manga.Attributes.Links["ap"])
				countHaveImagesExternal["ap"]++
			}
			if url != "" {
				data := []string{manga.Id, url}
				writeToCSV(writerAlternativeImage, data)
				countHaveImages++
			}

			// Debug
			if ct%100 == 0 {
				avgIterTime := float64(ct+1) / (1e-9 * float64(time.Since(start).Nanoseconds()))
				fmt.Printf("%d/%d mangas loaded at %.2f manga/sec (%d have images)....\n",
					ct, len(itemsManga), avgIterTime, countHaveImages)
			}

		}

		// Print out the number of covers we found
		fmt.Printf("done processing alternative images!\n")
		for key, value := range countHaveImagesExternal {
			fmt.Printf("\t %s had %d covers found\n", key, value)
		}

	}

}
