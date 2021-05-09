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
	"math"
	"net/http"
	"os"
	"time"
)

func downloadMangasBySearching(dirMangas string, dirMangasPrivate string, ctx context.Context, client *mangadex.APIClient,
	tagId2Tag *map[string]mangadex.TagResponse,
	mangasDownloaded *map[string]bool, tags []string, rating string) {

	// Cleaned tags (remove empty!)
	optsTags := make([]string, 0)
	for _, tag := range tags {
		if tag != "" {
			optsTags = append(optsTags, tag)
		}
	}

	// Specify our max limit and loop through the entire API to get all manga
	currentLimit := int32(100)
	maxOffset := int32(10000)
	for currentOffset := int32(0); currentOffset < maxOffset; currentOffset += currentLimit {

		// Perform our api search call to get the response
		opts := mangadex.MangaApiGetSearchMangaOpts{}
		opts.Limit = optional.NewInt32(currentLimit)
		opts.Offset = optional.NewInt32(currentOffset)
		if len(optsTags) != 0 {
			opts.IncludedTags = optional.NewInterface(optsTags)
		}
		if rating != "" {
			opts.ContentRating = optional.NewInterface(rating)
		}
		// robustly re-try a few times if we fail
		mangaList := mangadex.MangaList{}
		resp := &http.Response{}
		err := errors.New("startup")
		for retryCount := 0; retryCount <= 3 && err != nil; retryCount++ {
			mangaList, resp, err = client.MangaApi.GetSearchManga(ctx, &opts)
			if err != nil {
				fmt.Printf("\u001B[1;31mMANGA ERROR: %v\u001B[0m\n", err)
			} else if resp == nil {
				err = errors.New("invalid response object")
				fmt.Printf("\u001B[1;31mMANGA ERROR: respose object is nil\u001B[0m\n")
				continue
			} else if resp.StatusCode != 200 && resp.StatusCode != 204 {
				err = errors.New("invalid http error code")
				fmt.Printf("\u001B[1;31mMANGA ERROR: http code %d\u001B[0m\n", resp.StatusCode)
			}
			if err == nil {
				resp.Body.Close()
			}
			time.Sleep(250 * time.Millisecond)
		}

		// Debug print total for this tag
		if currentOffset == 0 {
			if rating != "" {
				fmt.Printf("%s | ", rating)
			}
			for _, tagId := range tags {
				if tagId == "" {
					fmt.Printf("EMPTY | ")
				} else {
					fmt.Printf("%s | ", (*tagId2Tag)[tagId].Data.Attributes.Name["en"])
				}
			}
			fmt.Printf("has %d total manga\n", mangaList.Total)
		}

		// Loop through all manga and print their ids
		for _, manga := range mangaList.Results {
			//fmt.Printf("%d/%d -> %s\n", currentOffset+int32(i), maxOffset, manga.Data.Id)
			if !(*mangasDownloaded)[manga.Data.Id] {
				file, _ := json.MarshalIndent(manga, "", " ")
				filePath := manga.Data.Id + ".json"
				if manga.Data.Attributes.ContentRating == "pornographic" ||
					manga.Data.Attributes.ContentRating == "erotica" {
					filePath = dirMangasPrivate + filePath
				} else {
					filePath = dirMangas + filePath
				}
				_ = ioutil.WriteFile(filePath, file, 0644)
				(*mangasDownloaded)[manga.Data.Id] = true
			}
		}

		// Update our current limit
		maxOffset = mangaList.Total
		currentLimit = int32(math.Min(float64(currentLimit), float64(maxOffset-currentOffset)))
		if currentOffset%200 == 0 || currentOffset+currentLimit >= maxOffset {
			fmt.Printf("\t - %d/%d completed....\n", currentOffset, maxOffset)
		}
		time.Sleep(250 * time.Millisecond)

	}

}

func main() {

	// Directory configuration
	dirMangas := "../data/manga/"
	dirMangasPrivate := "../data/manga_private/"
	dirChapters := "../data/chapter/"
	fileTagList := "../data/taglist.json"
	err := os.MkdirAll(dirMangas, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = os.MkdirAll(dirMangasPrivate, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = os.MkdirAll(dirChapters, os.ModePerm)
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

	// First get our tags
	tagList, resp, err := client.MangaApi.GetMangaTag(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("HTTP ERROR CODE %d\n", resp.StatusCode)
	}
	file, _ := json.MarshalIndent(tagList, "", " ")
	_ = ioutil.WriteFile(fileTagList, file, 0644)

	// Generate unique combinations we will try for our tags
	// This is to try to get as many mangas to be downloaded as possible
	// Since the api only returns 10k max, we will use tags to try to get all
	tagId2Tag := make(map[string]mangadex.TagResponse)
	tagIdList := make([]string, 0)
	for _, tag := range tagList {
		if tag.Data != nil && tag.Data.Type_ == "tag" {
			tagIdList = append(tagIdList, tag.Data.Id)
			tagId2Tag[tag.Data.Id] = tag
		}
	}
	tagIdList = append(tagIdList, "")
	tagIdListCombinations := make([][]string, 0)
	for _, tag1 := range tagIdList {
		for _, tag2 := range tagIdList {
			list := make([]string, 0)
			list = append(list, tag1)
			list = append(list, tag2)
			alreadyAdded := false
			for _, tags := range tagIdListCombinations {
				if tags[0] == list[1] && tags[1] == list[0] {
					alreadyAdded = true
					break
				}
			}
			if !alreadyAdded {
				tagIdListCombinations = append(tagIdListCombinations, list)
			}
		}
	}
	contentRatingIdList := []string{"safe", "suggestive", "erotica", "pornographic"}
	fmt.Printf("generated %d unique tag combinations from %d tags\n", len(tagIdListCombinations), len(tagIdList))

	// Here we will loop through all tags
	start := time.Now()
	mangasDownloaded := make(map[string]bool)
	for _, rating := range contentRatingIdList {
		downloadMangasBySearching(dirMangas, dirMangasPrivate, ctx, client, &tagId2Tag, &mangasDownloaded, []string{}, rating)
	}
	for _, tags := range tagIdListCombinations {
		downloadMangasBySearching(dirMangas, dirMangasPrivate, ctx, client, &tagId2Tag, &mangasDownloaded, tags, "")
	}
	fmt.Printf("downloaded %d mangas in %s!!\n\n", len(mangasDownloaded), time.Since(start))

}
