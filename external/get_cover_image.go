package external

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func GetCoverAniList(id string) string {

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
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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

	// Query MAL api for an image
	// https://jikan.docs.apiary.io/#reference/0/manga
	url := "https://api.jikan.moe/v3/manga/" + id + "/pictures"
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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

func GetCoverKitsu(id string) string {

	// Query kitsu api for an image
	// https://kitsu.docs.apiary.io/#reference/manga/manga/fetch-resource
	url := "https://kitsu.io/api/edge/manga?filter[slug]=" + id
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return ""
	}
	stringData, _ := ioutil.ReadAll(resp.Body)
	data := ResponseKitsuSearch{}
	err = json.Unmarshal(stringData, &data)
	if err != nil {
		return ""
	}
	if len(data.Data) < 1 {
		return ""
	}
	return data.Data[0].Attributes.Posterimage.Large

}

func GetCoverMangaUpdates(id string) string {

	// Query and get our html... (no api...)
	url := "https://www.mangaupdates.com/series.html?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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

	// Query and get our html... (no api...)
	url := "https://www.anime-planet.com/manga/" + id
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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
