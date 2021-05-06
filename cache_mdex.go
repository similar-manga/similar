package main

import (
	"./mangadex"
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)


func main() {

	// Directory configuration
	dirMangas := "data/manga/"
	dirChapters := "data/chapter/"

	// Create client
	config := mangadex.NewConfiguration()
	config.UserAgent = "similar-manga v2.0"
	config.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	client := mangadex.NewAPIClient(config)
	ctx := context.Background()

	// Specify our max limit and loop through the entire API to get all manga
	currentLimit := int32(100)
	maxOffset := int32(100000)
	for currentOffset := int32(0); currentOffset < maxOffset; currentOffset += currentLimit {

		// Perform our api search call to get the response
		opts := mangadex.MangaApiGetSearchMangaOpts{}
		opts.Limit = optional.NewInt32(currentLimit)
		opts.Offset = optional.NewInt32(currentOffset)
		mangaList, resp, err := client.MangaApi.GetSearchManga(ctx, &opts)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if resp.StatusCode != 200 {
			fmt.Println("HTTP ERROR CODE %d", resp.StatusCode)
			break
		}

		// Loop through all manga and print their ids
		for i, manga := range mangaList.Results {
			fmt.Printf("%d/%d -> %s\n", currentOffset+int32(i), maxOffset, manga.Data.Id)
			file, _ := json.MarshalIndent(manga.Data, "", " ")
			_ = ioutil.WriteFile(dirMangas+manga.Data.Id+".json", file, 0644)
		}

		// Update our current limit
		maxOffset = mangaList.Total
		currentLimit = int32(math.Min(float64(currentLimit), float64(maxOffset-currentOffset)))
		time.Sleep(200 * time.Millisecond)

	}

	// Loop through all manga and try to get their chapter information for each
	itemsManga, _ := ioutil.ReadDir(dirMangas)
	for _, file := range itemsManga {

		// Skip if a directory
		if file.IsDir() {
			continue
		}

		// Load the json from file into our manga struct
		manga := mangadex.Manga{}
		fileManga, _ := ioutil.ReadFile(dirMangas + file.Name())
		_ = json.Unmarshal([]byte(fileManga), &manga)

		// Perform our api search call to get the response
		opts := mangadex.MangaApiGetMangaIdFeedOpts{}
		opts.Limit = optional.NewInt32(500)
		chapterList, resp, err := client.MangaApi.GetMangaIdFeed(ctx, manga.Id, &opts)
		if resp != nil && resp.StatusCode == 404 {
			fmt.Printf("CHAPTER FEED GAVE %d (no chapter?!)\n", resp.StatusCode)
			continue
		}
		if err != nil {
			log.Fatalf("%v", err)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("HTTP ERROR CODE %d\n", resp.StatusCode)
			continue
		}

		// Loop through all chapter for this manga and save to disk
		fmt.Printf("manga %s -> %s\n", manga.Id, manga.Attributes.Title["en"])
		for _, chapter := range chapterList.Results {
			fmt.Printf("\t- chapter %s\n", chapter.Data.Id)
			fileChapter, _ := json.MarshalIndent(chapter.Data, "", " ")
			_ = ioutil.WriteFile(dirChapters+chapter.Data.Id+".json", fileChapter, 0644)
		}
		fmt.Println()
		time.Sleep(200 * time.Millisecond)

	}

}
