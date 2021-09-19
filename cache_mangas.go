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
	"strconv"
	"time"
)

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
	optsOrder := map[string]string{}
	optsOrder["createdAt"] = "asc"

	// Specify our max limit and loop through the entire API to get all manga
	currentLimit := int32(100)
	maxOffset := int32(10000)
	lastTimeApiCall := time.Now()
	for currentOffset := int32(0); currentOffset < maxOffset; currentOffset += currentLimit {

		// Perform our api search call to get the response
		opts := mangadex.MangaApiGetSearchMangaOpts{}
		opts.Limit = optional.NewInt32(currentLimit)
		opts.Offset = optional.NewInt32(currentOffset)
		opts.Order = optional.NewInterface(optsOrder)
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
		for retryCount := 0; retryCount <= 3 && err != nil; retryCount++ {

			// Rate limit if we have not waited enough
			minMilliBetween := int64(220)
			timeSinceLast := time.Since(lastTimeApiCall)
			if timeSinceLast.Milliseconds() < minMilliBetween {
				milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
				//fmt.Printf("\u001B[1;31mwaiting %d milliseconds\u001B[0m\n", milliToWait)
				time.Sleep(time.Duration(1e6 * milliToWait))
			}

			// Api call to the mangadex api (5 req per second)
			lastTimeApiCall = time.Now()
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
		maxOffset = int32(math.Min(float64(maxOffset), float64(mangaList.Total)))
		currentLimit = int32(math.Min(float64(currentLimit), float64(maxOffset-currentOffset)))
		if currentOffset%500 == 0 || currentOffset+currentLimit >= maxOffset {
			fmt.Printf("\t - %d/%d completed....\n", currentOffset, maxOffset)
		}

	}

}

func main() {

	// Directory configuration
	dirMangas := "../similar_data/manga/"
	fileTagList := "../similar_data/taglist.json"
	err := os.MkdirAll(dirMangas, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Create client
	config := mangadex.NewConfiguration()
	config.UserAgent = "similar-manga v2.1"
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
	mangasDownloaded := make(map[string]bool)
	contentRatingIdList := []string{"safe", "suggestive", "erotica", "pornographic"}
	for _, rating := range contentRatingIdList {
		downloadMangasBySearching(dirMangas, ctx, client, &tagId2Tag, &mangasDownloaded, []string{}, rating, "")
	}
	for _, tags := range tagIdListCombinations {
		downloadMangasBySearching(dirMangas, ctx, client, &tagId2Tag, &mangasDownloaded, tags, "safe", "")
	}
	for year := 2018; year <= 2021; year++ {
		for month := 1; month <= 12; month++ {
			createdAtSince := strconv.Itoa(year) + "-" + fmt.Sprintf("%02d", month) + "-01T00:00:00"
			downloadMangasBySearching(dirMangas, ctx, client, &tagId2Tag, &mangasDownloaded, []string{}, "safe", createdAtSince)
		}
	}
	fmt.Printf("downloaded %d mangas in %s!!\n\n", len(mangasDownloaded), time.Since(start))

}
