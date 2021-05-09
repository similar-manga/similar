package main

import (
	"./mangadex"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Directory configuration
	dirMangas := "../data/manga/"
	dirChapters := "../data/chapter/"
	err := os.MkdirAll(dirChapters, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Create client
	config := mangadex.NewConfiguration()
	config.UserAgent = "similar-manga v2.0"
	config.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	client := mangadex.NewAPIClient(config)
	ctx := context.Background()

	// Loop through all manga and try to get their chapter information for each
	countChaptersDownloaded := 0
	start := time.Now()
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

		// Perform our api search call to get the response
		opts := mangadex.MangaApiGetMangaIdFeedOpts{}
		opts.Limit = optional.NewInt32(500)

		// robustly re-try a few times if we fail
		chapterList := mangadex.ChapterList{}
		resp := &http.Response{}
		err := errors.New("startup")
		for retryCount := 0; retryCount <= 3 && err != nil; retryCount++ {
			chapterList, resp, err = client.MangaApi.GetMangaIdFeed(ctx, manga.Data.Id, &opts)
			if err != nil {
				fmt.Printf("\u001B[1;31mCHAPTER ERROR: %v\u001B[0m\n", err)
			} else if resp == nil {
				err = errors.New("invalid response object")
				fmt.Printf("\u001B[1;31mCHAPTER ERROR: respose object is nil\u001B[0m\n")
				continue
			} else if resp.StatusCode != 200 && resp.StatusCode != 204 && resp.StatusCode != 404 {
				err = errors.New("invalid http error code")
				fmt.Printf("\u001B[1;31mCHAPTER ERROR: http code %d\u001B[0m\n", resp.StatusCode)
			}
			if err == nil {
				resp.Body.Close()
			}
			time.Sleep(250 * time.Millisecond)
		}

		// Loop through all chapter for this manga and save to disk
		fmt.Printf("manga %s -> %s\n", manga.Data.Id, manga.Data.Attributes.Title["en"])
		for _, chapter := range chapterList.Results {
			fmt.Printf("\t- chapter %s\n", chapter.Data.Id)
			fileChapter, _ := json.MarshalIndent(chapter, "", " ")
			_ = ioutil.WriteFile(dirChapters+chapter.Data.Id+".json", fileChapter, 0644)
			countChaptersDownloaded++
		}
		fmt.Println()

	}
	fmt.Printf("downloaded %d chapters in %s!!\n\n", countChaptersDownloaded, time.Since(start))

}
