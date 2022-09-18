package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antihax/optional"
	"github.com/similar-manga/similar/mangadex"
	"github.com/similar-manga/similar/similar"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	// Directory configuration
	// 39147 mangas in total as of 03-19-2022
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
	dirMangas := dirData + "manga/"
	dirChapters := dirData + "chapter/"
	dirChaptersInfo := dirData + "chapter_info/"

	// Check that we have valid range
	if algoNumMin != -1 && algoNumMin >= algoNumMax {
		log.Fatalf("invalid range of %d to %d\n", algoNumMin, algoNumMax)
	}

	// Settings
	skipAlreadyDownloaded := true
	saveRawChapterList := false
	if saveRawChapterList {
		err := os.MkdirAll(dirChapters, os.ModePerm)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
	err := os.MkdirAll(dirChaptersInfo, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Create client
	config := mangadex.NewConfiguration()
	config.UserAgent = "similar-manga v2.2"
	config.HTTPClient = &http.Client{
		Timeout: 60 * time.Second,
	}
	client := mangadex.NewAPIClient(config)
	ctx := context.Background()

	// Loop through all manga and try to get their chapter information for each
	countMangasDownloaded := 0
	countChaptersDownloaded := 0
	start := time.Now()
	itemsManga, _ := ioutil.ReadDir(dirMangas)
	lastTimeApiCall := time.Now()
	fmt.Printf("starting download of chapters for %d mangas\n", len(itemsManga))
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
		_ = json.Unmarshal(fileManga, &manga)

		// Either try to re-download or download if we don't have the chapter
		chapterFilePath := dirChapters + manga.Id + ".json"
		_, err := os.Stat(chapterFilePath)
		chapterList := mangadex.ChapterList{}
		if !skipAlreadyDownloaded || os.IsNotExist(err) {

			// Default includes we should use!
			optsIncludes := make([]string, 0)
			optsIncludes = append(optsIncludes, "user")
			optsIncludes = append(optsIncludes, "scanlation_group")

			// Perform our api search call to get the response
			opts := mangadex.MangaApiGetMangaIdFeedOpts{}
			opts.Limit = optional.NewInt32(100)
			opts.Offset = optional.NewInt32(0)
			opts.Includes = optional.NewInterface(optsIncludes)

			// Robustly re-try a few times if we fail
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
				chapterList, resp, err = client.MangaApi.GetMangaIdFeed(ctx, manga.Id, &opts)
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

			}

			// Write chapter this for this manga to file
			if saveRawChapterList {
				file, _ := json.MarshalIndent(chapterList, "", " ")
				_ = ioutil.WriteFile(chapterFilePath, file, 0644)
			}
			countChaptersDownloaded += len(chapterList.Data)

		} else {
			// check that we have the chapter
			if os.IsNotExist(err) {
				fmt.Printf("%d/%d mangas -> manga chapters now downloaded!!!\n", ct+1)
				continue
			}
			// Now lets open the file and do some computations
			chapterList = mangadex.ChapterList{}
			fileChapter, _ := ioutil.ReadFile(chapterFilePath)
			_ = json.Unmarshal(fileChapter, &chapterList)
		}

		// Get compress "information" about this chapter such as the number of chapters
		// Languages, and what scanlation groups have translated for this
		chapterInfo := similar.ChapterInformation{}
		chapterInfo.Id = manga.Id
		chapterInfo.NumChapters = len(chapterList.Data)
		tempLanguages := map[string]bool{}
		tempGroups := map[string]similar.ChapterGroup{}
		for _, chapter := range chapterList.Data {
			lang := chapter.Attributes.TranslatedLanguage
			group := similar.ChapterGroup{Id: "unknown", Name: "unknown"}
			for _, relation := range chapter.Relationships {
				if relation.Type_ == "scanlation_group" && relation.Attributes != nil {
					attributes := (*relation.Attributes).(map[string]interface{})
					group = similar.ChapterGroup{Id: relation.Id, Name: attributes["name"].(string)}
					break
				}
			}
			// Append to our maps if not added
			if _, ok := tempLanguages[lang]; !ok {
				tempLanguages[lang] = true
			}
			if _, ok := tempGroups[group.Id]; !ok && group.Id != "unknown" {
				tempGroups[group.Id] = group
			}
		}
		for k := range tempLanguages {
			chapterInfo.Languages = append(chapterInfo.Languages, k)
		}
		for _, v := range tempGroups {
			chapterInfo.Groups = append(chapterInfo.Groups, v)
		}

		// Finally write the info to file
		chapterInfoFilePath := dirChaptersInfo + manga.Id + ".json"
		file, _ := json.MarshalIndent(chapterInfo, "", " ")
		_ = ioutil.WriteFile(chapterInfoFilePath, file, 0644)

		// Debug print
		if (ct+1)%200 == 0 {
			avgIterTime := float64(ct+1) / time.Since(start).Seconds()
			totalMangas := len(itemsManga)
			if algoNumMin != -1 {
				totalMangas = algoNumMax - algoNumMin
			}
			fmt.Printf("%d/%d mangas -> %d chapter downloaded at %.2f manga/sec....\n", ct+1, totalMangas, countChaptersDownloaded, avgIterTime)
		}
		countMangasDownloaded++

	}
	fmt.Printf("downloaded %d chapters from %d mangas in %s!!\n\n", countChaptersDownloaded, countMangasDownloaded, time.Since(start))

}
