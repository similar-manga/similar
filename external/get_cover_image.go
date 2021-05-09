package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"time"
)

var lastTimeCoverKT = time.Now()
var lastTimeCoverAL = time.Now()
var lastTimeCoverMAL = time.Now()
var lastTimeCoverMU = time.Now()
var lastTimeCoverAP = time.Now()

func GetCoverKitsu(id string) string {

	// Rate limit if we have not waited enough
	minMilliBetween := int64(50)
	timeSinceLast := time.Since(lastTimeCoverKT)
	if timeSinceLast.Milliseconds() < minMilliBetween {
		milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
		//fmt.Printf("\u001B[1;31mEXTERNAL KT: waiting %d milliseconds\u001B[0m\n", milliToWait)
		time.Sleep(time.Duration(1e6 * milliToWait))
	}

	// Query kitsu api for an image
	// https://kitsu.docs.apiary.io/#reference/manga/manga/fetch-resource
	//url := "https://kitsu.io/api/edge/manga?filter[slug]=" + id
	//resp, err := http.Get(url)
	//lastTimeCoverKT = time.Now()
	//time.Sleep(500 * time.Millisecond)
	//if err != nil {
	//	return ""
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != 200 {
	//	fmt.Printf("\u001B[1;31mEXTERNAL KT: http code %d\u001B[0m\n", resp.StatusCode)
	//	if resp.StatusCode == 429 {
	//		time.Sleep(time.Second)
	//	}
	//	return ""
	//}
	//stringData, _ := ioutil.ReadAll(resp.Body)
	//data := ResponseKitsuSearch{}
	//err = json.Unmarshal(stringData, &data)
	//if err != nil {
	//	return ""
	//}
	//if len(data.Data) < 1 {
	//	return ""
	//}
	//return data.Data[0].Attributes.Posterimage.Large

	// Construct the kistu image url
	url := "https://media.kitsu.io/manga/poster_images/" + id + "/large.jpg"
	resp, err := http.Get(url)
	lastTimeCoverKT = time.Now()
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("\u001B[1;31mEXTERNAL KT: http code %d\u001B[0m\n", resp.StatusCode)
		if resp.StatusCode == 429 {
			time.Sleep(time.Second)
		}
		return ""
	}
	return url

}

func GetCoverAniList(id string) string {

	// Rate limit if we have not waited enough
	minMilliBetween := int64(700)
	timeSinceLast := time.Since(lastTimeCoverAL)
	if timeSinceLast.Milliseconds() < minMilliBetween {
		milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
		//fmt.Printf("\u001B[1;31mEXTERNAL AL: waiting %d milliseconds\u001B[0m\n", milliToWait)
		time.Sleep(time.Duration(1e6 * milliToWait))
	}

	// Query graph ql endpoint for our image
	// https://anilist.gitbook.io/anilist-apiv2-docs/overview/graphql/getting-started
	jsonData := map[string]string{
		"query": `
            {
			  Media(id: ` + id + `, type: MANGA) {
				coverImage {
				  extraLarge
				}
			  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(jsonValue))
	lastTimeCoverAL = time.Now()
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("\u001B[1;31mEXTERNAL AL: http code %d\u001B[0m\n", resp.StatusCode)
		if resp.StatusCode == 429 {
			time.Sleep(time.Second)
		}
		return ""
	}
	stringData, _ := ioutil.ReadAll(resp.Body)
	data := ResponseAniList{}
	err = json.Unmarshal(stringData, &data)
	if err != nil {
		return ""
	}
	return data.Data.Media.Coverimage.Extralarge

}

func GetCoverMyAnimeList(id string) string {

	// Rate limit if we have not waited enough
	minMilliBetween := int64(50)
	timeSinceLast := time.Since(lastTimeCoverMAL)
	if timeSinceLast.Milliseconds() < minMilliBetween {
		milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
		//fmt.Printf("\u001B[1;31mEXTERNAL MAL: waiting %d milliseconds\u001B[0m\n", milliToWait)
		time.Sleep(time.Duration(1e6 * milliToWait))
	}

	// Query MAL api for an image
	// https://jikan.docs.apiary.io/#reference/0/manga
	url := "https://api.jikan.moe/v3/manga/" + id + "/pictures"
	resp, err := http.Get(url)
	lastTimeCoverMAL = time.Now()
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("\u001B[1;31mEXTERNAL MAL: http code %d\u001B[0m\n", resp.StatusCode)
		if resp.StatusCode == 429 {
			time.Sleep(time.Second)
		}
		return ""
	}
	stringData, _ := ioutil.ReadAll(resp.Body)
	data := ResponseAnimeList{}
	err = json.Unmarshal(stringData, &data)
	if err != nil {
		return ""
	}
	if len(data.Pictures) < 1 {
		return ""
	}
	return data.Pictures[0].Large

}

func GetCoverMangaUpdates(id string) string {

	// Rate limit if we have not waited enough
	minMilliBetween := int64(50)
	timeSinceLast := time.Since(lastTimeCoverMU)
	if timeSinceLast.Milliseconds() < minMilliBetween {
		milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
		//fmt.Printf("\u001B[1;31mEXTERNAL MU: waiting %d milliseconds\u001B[0m\n", milliToWait)
		time.Sleep(time.Duration(1e6 * milliToWait))
	}

	// Query and get our html... (no api...)
	url := "https://www.mangaupdates.com/series.html?id=" + id
	resp, err := http.Get(url)
	lastTimeCoverMU = time.Now()
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("\u001B[1;31mEXTERNAL MU: http code %d\u001B[0m\n", resp.StatusCode)
		if resp.StatusCode == 429 {
			time.Sleep(time.Second)
		}
		return ""
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ""
	}

	// Return the url of the image
	// Logic found using google chrome (right click in inspector and copy "selector")
	return doc.Find("#main_content > div > div.row.no-gutters "+
		"> div > div > center > img").AttrOr("src", "")

}

func GetCoverAnimePlanet(id string) string {

	// Rate limit if we have not waited enough
	minMilliBetween := int64(50)
	timeSinceLast := time.Since(lastTimeCoverAP)
	if timeSinceLast.Milliseconds() < minMilliBetween {
		milliToWait := minMilliBetween - timeSinceLast.Milliseconds()
		//fmt.Printf("\u001B[1;31mEXTERNAL AP: waiting %d milliseconds\u001B[0m\n", milliToWait)
		time.Sleep(time.Duration(1e6 * milliToWait))
	}

	// Query and get our html... (no api...)
	url := "https://www.anime-planet.com/manga/" + id
	resp, err := http.Get(url)
	lastTimeCoverAP = time.Now()
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("\u001B[1;31mEXTERNAL AP: http code %d\u001B[0m\n", resp.StatusCode)
		if resp.StatusCode == 429 {
			time.Sleep(time.Second)
		}
		return ""
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ""
	}

	// Return the url of the image
	// Logic found using google chrome (right click in inspector and copy "selector")
	return "https://www.anime-planet.com" +
		doc.Find("#entry > div > div > div > div > img").AttrOr("src", "")

}
