package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antihax/optional"
	"github.com/similar-manga/similar/mangadex"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

// This is a global variable so we rate limit over all API calls
var lastTimeApiCall = time.Now()

func downloadMangasBySearching(dirMangas string, ctx context.Context, client *mangadex.APIClient,
	tagId2Tag *map[string]mangadex.Tag,
	mangasDownloaded *map[string]bool, tags []string, rating string, createdAtSince string) {

	// Cleaned tags (remove empty!)
	optsTags := make([]string, 0)
	for _, tag := range tags {
		if tag != "" {
			optsTags = append(optsTags, tag)
		}
	}

	// Default includes we should use!
	optsIncludes := make([]string, 0)
	optsIncludes = append(optsIncludes, "author")
	optsIncludes = append(optsIncludes, "artist")
	optsIncludes = append(optsIncludes, "cover_art")

	// Specify our max limit and loop through the entire API to get all manga
	currentLimit := int32(100)
	maxOffset := int32(10000)
	for currentOffset := int32(0); currentOffset < maxOffset; currentOffset += currentLimit {

		// Perform our api search call to get the response
		opts := mangadex.MangaApiGetSearchMangaOpts{}
		opts.Limit = optional.NewInt32(currentLimit)
		opts.Offset = optional.NewInt32(currentOffset)
		opts.OrderCreatedAt = optional.NewString("asc")
		if len(optsTags) != 0 {
			opts.IncludedTags = optional.NewInterface(optsTags)
		}
		if rating != "" {
			opts.ContentRating = optional.NewInterface(rating)
		}
		if createdAtSince != "" {
			opts.CreatedAtSince = optional.NewString(createdAtSince)
		}
		opts.Includes = optional.NewInterface(optsIncludes)

		// robustly re-try a few times if we fail
		mangaList := mangadex.MangaList{}
		resp := &http.Response{}
		err := errors.New("startup")
		maxRetries := 10
		for retryCount := 0; retryCount <= maxRetries && err != nil; retryCount++ {

			// Pause if we need to retry as we are probably rate limited...
			if retryCount > 0 {
				//fmt.Printf("\u001B[1;31mretrying %d / %d times...\u001B[0m\n", retryCount, maxRetries)
				time.Sleep(5 * time.Second)
			}

			// Rate limit if we have not waited enough
			// NOTE: /manga has 10 reqs per 60 minutes limit (seems really slow...)
			minMilliBetween := int64(500)
			timeSinceLast := time.Since(lastTimeApiCall)
			if timeSinceLast.Milliseconds() < minMilliBetween {
				milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
				//fmt.Printf("\u001B[1;31mwaiting %d milliseconds\u001B[0m\n", milliToWait)
				time.Sleep(time.Duration(milliToWait) * time.Millisecond)
			}

			// Api call to the mangadex api (5 req per second)
			lastTimeApiCall = time.Now()
			mangaList, resp, err = client.MangaApi.GetSearchManga(ctx, &opts)
			if err != nil {
				fmt.Printf("\u001B[1;31mMANGA ERROR (%d of %d): %v\u001B[0m\n", retryCount, maxRetries, err)
			} else if resp == nil {
				err = errors.New("invalid response object")
				fmt.Printf("\u001B[1;31mMANGA ERROR (%d of %d): respose object is nil\u001B[0m\n", retryCount, maxRetries)
				continue
			} else if resp.StatusCode != 200 && resp.StatusCode != 204 {
				err = errors.New("invalid http error code")
				fmt.Printf("\u001B[1;31mMANGA ERROR (%d of %d): http code %d\u001B[0m\n", retryCount, maxRetries, resp.StatusCode)
			}
			if err == nil {
				resp.Body.Close()
			}

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
					tmpTag := (*tagId2Tag)[tagId]
					tmpName := (*tmpTag.Attributes).Name
					fmt.Printf("%s | ", (*tmpName)["en"])
				}
			}
			if createdAtSince != "" {
				fmt.Printf("%s | \n", createdAtSince)
			}
			fmt.Printf("has %d total manga\n", mangaList.Total)
		}

		// Loop through all manga and print their ids
		for _, manga := range mangaList.Data {
			//fmt.Printf("%d/%d -> %s\n", currentOffset+int32(i), maxOffset, manga.Data.Id)
			//fmt.Printf("manga - %s \n", manga.Attributes.CreatedAt)
			if !(*mangasDownloaded)[manga.Id] {
				file, _ := json.MarshalIndent(manga, "", " ")
				_ = ioutil.WriteFile(dirMangas+manga.Id+".json", file, 0644)
				(*mangasDownloaded)[manga.Id] = true
			}
		}

		// Update our current limit
		// NOTE: they have coded a hard max of 10k, thus don't use the reported one...
		if mangaList.Total > 0 {
			maxOffset = int32(math.Min(float64(maxOffset), float64(mangaList.Total)))
		}
		currentLimit = int32(math.Min(float64(currentLimit), float64(maxOffset-currentOffset)))
		if currentOffset%500 == 0 || currentOffset+currentLimit >= maxOffset {
			fmt.Printf("\t - %d/%d completed....\n", currentOffset, maxOffset)
		}

	}

}

func main() {

	// Directory configuration
	// Algorithm number:
	// 		-1 - try to query all mangas
	// 		0-45: select which query to perform
	dirData := "D:/MANGADEX/similar_data/"
	algoNum := -1
	if len(os.Args) == 2 {
		dirData = os.Args[1]
	}
	if len(os.Args) == 3 {
		dirData = os.Args[1]
		algoNum, _ = strconv.Atoi(os.Args[2])
	}
	dirMangas := dirData + "manga/"
	fileTagList := dirData + "taglist.json"
	err := os.MkdirAll(dirMangas, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("directory %s\n", dirData)
	fmt.Printf("  - mangas %s\n", dirMangas)
	fmt.Printf("  - tags %s\n", fileTagList)

	// Create client
	config := mangadex.NewConfiguration()
	config.UserAgent = "similar-manga v2.3"
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
	tagId2Tag := make(map[string]mangadex.Tag)
	tagIdList := make([]string, 0)
	for _, tag := range tagList.Data {
		if tag.Attributes != nil && tag.Type_ == "tag" {
			tagIdList = append(tagIdList, tag.Id)
			tagId2Tag[tag.Id] = tag
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
	fmt.Printf("generated %d unique tag combinations from %d tags\n", len(tagIdListCombinations), len(tagIdList))

	// Here we will loop through all tags
	start := time.Now()
	algoCurr := 1
	mangasDownloaded := make(map[string]bool)
	contentRatingIdList := []string{"safe", "suggestive", "erotica", "pornographic"}
	if algoNum == -1 || algoNum == 1 {
		for _, rating := range contentRatingIdList {
			downloadMangasBySearching(dirMangas, ctx, client, &tagId2Tag, &mangasDownloaded, []string{}, rating, "")
		}
	}
	algoCurr++
	for ct, tags := range tagIdListCombinations {
		if algoNum == -1 || algoNum == algoCurr {
			downloadMangasBySearching(dirMangas, ctx, client, &tagId2Tag, &mangasDownloaded, tags, "safe", "")
		}
		if (ct+1)%len(tagIdList) == 0 {
			algoCurr++
		}
	}
	for year := 2018; year <= time.Now().Year(); year++ {
		if algoNum == -1 || algoNum == algoCurr {
			for month := 1; month <= 12; month++ {
				createdAtSince := strconv.Itoa(year) + "-" + fmt.Sprintf("%02d", month) + "-01T00:00:00"
				downloadMangasBySearching(dirMangas, ctx, client, &tagId2Tag, &mangasDownloaded, []string{}, "safe", createdAtSince)
			}
		}
		algoCurr++
	}
	fmt.Printf("processed of %d call of %d total\n", algoNum, algoCurr)
	fmt.Printf("downloaded %d mangas in %s!!\n\n", len(mangasDownloaded), time.Since(start))

}
