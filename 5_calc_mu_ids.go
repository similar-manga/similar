package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/similar-manga/similar/external"
	"github.com/similar-manga/similar/mangadex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Directory configuration
	dirMangas := "D:/MANGADEX/similar_data/manga/"
	dirMappings := "D:/MANGADEX/similar_data/mapping/"
	minMilliBetween := int64(600)
	err := os.MkdirAll(dirMappings, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// For our ID conversion
	// https://www.unitconverters.net/numbers/base-36-to-decimal.htm
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	// mangaupdates
	// https://www.mangaupdates.com/series.html?id=`{id}`
	// https://api.mangaupdates.com/#operation/retrieveSeries
	// https://api.mangaupdates.com/v1/series/(base38 encoding of 7char ids)
	// https://api.mangaupdates.com/v1/series/66788345008/rss
	fileMU := external.OpenCSVFileStream(dirMappings + "mangaupdates_new2mdex.csv")
	defer fileMU.Close()
	writerMU := csv.NewWriter(fileMU)
	defer writerMU.Flush()

	// Loop through all manga and try to get their chapter information for each
	start := time.Now()
	lastTime := time.Now()
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
		if _, ok := manga.Attributes.Links["mu"]; ok {
			muLink := manga.Attributes.Links["mu"]
			if len(muLink) == 7 {

				// Encode from base36 format
				idEncoded := int64(external.Decode(muLink))
				//fmt.Printf("%d/%d manga %s ->%s encoded into %s\n", ct+1, len(itemsManga), manga.Id, muLink, strconv.FormatInt(idEncoded, 10))

				// Try the new id!
				resp2, err2 := http.Get("https://api.mangaupdates.com/v1/series/" + strconv.FormatInt(idEncoded, 10))
				if err2 == nil {
					defer resp2.Body.Close()
				}

				// Save if good!
				if err2 == nil && resp2.StatusCode == 200 {
					fmt.Printf("%d/%d manga %s -> mu id %s encoded into %s -> is new MU id!\n", ct+1, len(itemsManga), manga.Id, muLink, strconv.FormatInt(idEncoded, 10))
					data := []string{strconv.FormatInt(idEncoded, 10), manga.Id}
					external.WriteToCSV(writerMU, data)
				}

			} else if re.MatchString(muLink) {
				ints := re.FindAllString(muLink, -1)
				idOriginal, err := strconv.Atoi(ints[0])
				if err == nil {

					// Rate limit if we have not waited enough
					timeSinceLast := time.Since(lastTime)
					if timeSinceLast.Milliseconds() < minMilliBetween {
						milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
						time.Sleep(time.Duration(1e6 * milliToWait))
					}

					// Try the existing as the id (not likely since mangadex won't have updated..)
					resp1, err1 := http.Get("https://api.mangaupdates.com/v1/series/" + strconv.Itoa(idOriginal))
					if err1 == nil {
						defer resp1.Body.Close()
					}
					lastTime = time.Now()

					// debug print it out
					if err1 == nil && resp1.StatusCode == 200 {
						fmt.Printf("%d/%d manga %s -> mu id of %d -> is old MU id...\n", ct+1, len(itemsManga), manga.Id, idOriginal)
						data := []string{strconv.Itoa(idOriginal), manga.Id}
						external.WriteToCSV(writerMU, data)
					} else {

						// We have a couple retires here
						ctr := 0
						ctrMax := 5
						found := false
						for ctr < ctrMax {

							// Rate limit if we have not waited enough
							timeSinceLast = time.Since(lastTime)
							if timeSinceLast.Milliseconds() < minMilliBetween {
								milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
								time.Sleep(time.Duration(1e6 * milliToWait))
							}

							// If invalid, then try to get the page and parse it!
							// Query and get our html... (no api to get this...)
							url := "https://www.mangaupdates.com/series.html?id=" + strconv.Itoa(idOriginal)
							resp, err := http.Get(url)
							lastTime = time.Now()
							if err == nil {
								defer resp.Body.Close()
							}

							// Sleep if we get a warning, otherwise we don't retry again!
							if err == nil && resp.StatusCode == 429 {
								fmt.Printf("\u001B[1;31mEXTERNAL MU: http code %d (try %d of %d)\u001B[0m\n", resp.StatusCode, ctr, ctrMax)
								time.Sleep(2.0 * time.Second)
							}
							if err == nil && resp.StatusCode != 429 {
								ctr = ctrMax
							}

							// Load the HTML document
							// Logic found using google chrome (right click in inspector and copy "selector")
							if err == nil && resp.StatusCode == 200 {
								doc, err := goquery.NewDocumentFromReader(resp.Body)
								if err == nil {
									rssUrl := doc.Find("#main_content > div:nth-child(2) > div.row.no-gutters > div.col-12.p-2 > a").AttrOr("href", "")
									paths := strings.Split(rssUrl, "/")
									if len(paths) > 3 {
										rssId := paths[len(paths)-2]
										fmt.Printf("%d/%d manga %s -> mu id of %d | RSS URL IS %s | %s id found\n", ct+1, len(itemsManga), manga.Id, idOriginal, rssUrl, rssId)
										data := []string{rssId, manga.Id}
										external.WriteToCSV(writerMU, data)
										writerMU.Flush()
										found = true
									}
								}
							}
							ctr += 1
						}
						if !found {
							fmt.Printf("%d/%d manga %s -> mu invalid %s\n", ct+1, len(itemsManga), manga.Id, muLink)
						}
					}
				}

			}
		}

		// Debug
		//if ct%100 == 0 {
		//	avgIterTime := float64(ct+1) / time.Since(start).Seconds()
		//	fmt.Printf("%d/%d mangas -> processing at %.2f manga/sec....\n", ct+1, len(itemsManga), avgIterTime)
		//}

	}
	writerMU.Flush()
	fileMU.Close()
	fmt.Printf("done processing mappings (%.2f seconds)!\n", time.Since(start).Seconds())

}
