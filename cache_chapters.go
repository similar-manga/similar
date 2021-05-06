package main

import (
	"./mangadex"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func reportToMangadexNetwork(url string, filename string, start time.Time, success bool, cached bool) {

	// Create default
	values := make(map[string]interface{})
	values["url"] = url
	values["success"] = success
	values["bytes"] = 0
	values["duration"] = time.Since(start).Milliseconds()
	values["cached"] = cached

	// If failed directly report
	if !success {
		jsonValue, _ := json.Marshal(values)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("MD@HOME: %v", err)
		} else {
			resp.Body.Close()
		}
		return
	}

	// If file does not exists then we have already failed
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		values["success"] = false
		jsonValue, _ := json.Marshal(values)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("MD@HOME: %v", err)
		} else {
			resp.Body.Close()
		}
		return
	}

	// Finally report the downloaded image to mangadex @ home network report
	fi, _ := os.Stat(filename)
	values["bytes"] = fi.Size()
	jsonValue, _ := json.Marshal(values)
	//fmt.Println(string(jsonValue))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("MD@HOME: %v", err)
	} else {
		resp.Body.Close()
	}

}

func downloadChapterImage(chapterPath string, chapter mangadex.Chapter, image string, baseUrl string) {

	// Create the url we will download
	start := time.Now()
	filename := chapterPath + image
	url := baseUrl + "/data/" + chapter.Attributes.Hash + "/" + image
	//fmt.Printf("%d/%d (image %d/%d) -> %s\n", i, totalChapters, c+1, len(chapter.Attributes.Data), url)

	// Skip if already downloaded
	if _, err := os.Stat(filename); err == nil {
		return
	}

	// Try to download
	imgResp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v\n", err)
		reportToMangadexNetwork(url, filename, start, false, false)
		return
	}
	cacheHit := imgResp.Header.Get("X-Cache")

	// Open a file for writing and write the response!
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
		reportToMangadexNetwork(url, filename, start, false, cacheHit == "HIT")
		return
	}
	_, err = io.Copy(file, imgResp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		reportToMangadexNetwork(url, filename, start, false, cacheHit == "HIT")
		return
	}
	imgResp.Body.Close()
	file.Close()

	// Report to mangadex @ home network!
	reportToMangadexNetwork(url, filename, start, true, cacheHit == "HIT")

}

func main() {

	// Directory configuration
	dirChapters := "data/chapter/"
	dirImages := "data/images/"

	// Create client
	config := mangadex.NewConfiguration()
	client := mangadex.NewAPIClient(config)
	config.UserAgent = "similar-manga v2.0"
	config.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	ctx := context.Background()

	// Loop through all manga and download each chapter's images!
	itemsChapters, _ := ioutil.ReadDir(dirChapters)
	for _, file := range itemsChapters {

		// Skip if a directory
		if file.IsDir() {
			continue
		}

		// Load the json from file into our chapter struct
		chapter := mangadex.Chapter{}
		fileManga, _ := ioutil.ReadFile(dirChapters + file.Name())
		_ = json.Unmarshal([]byte(fileManga), &chapter)

		// Skip if not in english
		if chapter.Attributes.TranslatedLanguage != "en" {
			continue
		}

		// Create our save folder path
		fmt.Printf("chapter %s\n", chapter.Id)
		chapterPath := dirImages + chapter.Id + "/"
		err := os.MkdirAll(chapterPath, os.ModePerm)
		if err != nil {
			log.Fatalf("%v", err)
		}

		// Get the mangadex@home url we will download the images from
		opts := mangadex.AtHomeApiGetAtHomeServerChapterIdOpts{}
		mdexAtHome, resp, err := client.AtHomeApi.GetAtHomeServerChapterId(ctx, chapter.Id, &opts)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("HTTP ERROR CODE %d\n", resp.StatusCode)
			continue
		}

		// Create our worker pool which will try to download many chapters
		start := time.Now()
		var wg sync.WaitGroup
		workerPoolSize := 10
		dataCh := make(chan string, workerPoolSize)
		for w := 0; w < workerPoolSize; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for data := range dataCh {
					downloadChapterImage(chapterPath, chapter, data, mdexAtHome.BaseUrl)
					fmt.Printf("\t- downloaded %s\n", data)
				}
			}()
		}

		// Now feed data into our channel till it is done
		for _, image := range chapter.Attributes.Data {
			dataCh <- image
		}
		close(dataCh)
		wg.Wait()
		fmt.Println()
		fmt.Printf("chapter took %s\n", time.Since(start))
		time.Sleep(200 * time.Millisecond)

	}

}
