package similar

import (
	"github.com/similar-manga/similar/mangadex"
	"strings"
)

var OneWayTags = []mangadex.Tag{
	{
		Id:    "b11fda93-8f1d-4bef-b2ed-8803d3733170",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "4-Koma"},
			Version: 1,
		},
	},
	{
		Id:    "b13b2a48-c720-44a9-9c77-39c9979373fb",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Doujinshi"},
			Version: 1,
		},
	},
	{
		Id:    "b29d6a3d-1569-4e7a-8caf-7557bc92cd5d",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Gore"},
			Version: 1,
		},
	},
	{
		Id:    "97893a4c-12af-4dac-b6be-0dffb353568e",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Sexual Violence"},
			Version: 1,
		},
	},
	{
		Id:    "5920b825-4181-4a17-beeb-9918b0ff7a30",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Boys' Love"},
			Version: 1,
		},
	},
	{
		Id:    "a3c67850-4684-404e-9b7f-c69850ee5da6",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Girls' Love"},
			Version: 1,
		},
	},
	{
		Id:    "acc803a4-c95a-4c22-86fc-eb6b582d82a2",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Wuxia"},
			Version: 1,
		},
	},
	{
		Id:    "2d1f5d56-a1e5-4d0d-a961-2193588b08ec",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Loli"},
			Version: 1,
		},
	},
	{
		Id:    "ddefd648-5140-4e5f-ba18-4eca4071d19b",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Shota"},
			Version: 1,
		},
	},
	{
		Id:    "5bd0e105-4481-44ca-b6e7-7544da56b1a3",
		Type_: "tag",
		Attributes: &mangadex.TagAttributes{
			Name:    &map[string]string{"en": "Incest"},
			Version: 1,
		},
	},
}

func NotValidMatch(manga mangadex.Manga, mangaOther mangadex.Manga) bool {

	// Enforce that our two demographics are the same
	if manga.Attributes.ContentRating != "" &&
		manga.Attributes.ContentRating != mangaOther.Attributes.ContentRating {
		return true
	}

	// Enforce that our two demographics are the same
	if manga.Attributes.PublicationDemographic != "" &&
		manga.Attributes.PublicationDemographic != mangaOther.Attributes.PublicationDemographic {
		return true
	}

	// No need to check tags for our top level content ratings
	// They will be a valid match no matter the tags (not that many options thus can't limit)
	if manga.Attributes.ContentRating == "erotica" || manga.Attributes.ContentRating == "pornographic" {
		return false
	}

	// Enforce that the two do not have another as a *related* manga
	// A related manga is a sidestory / sequal / prequel for the work
	related := map[string]bool{}
	relatedOther := map[string]bool{}
	for _, relation := range manga.Relationships {
		if relation.Type_ == "manga" {
			related[relation.Id] = true
		}
	}
	for _, relation := range mangaOther.Relationships {
		if relation.Type_ == "manga" {
			relatedOther[relation.Id] = true
		}
	}
	if _, found := related[mangaOther.Id]; found {
		return false
	}
	if _, found := relatedOther[manga.Id]; found {
		return false
	}


	// Next we should enforce the following tags
	for _, tag1 := range OneWayTags {

		// Check to see if this tag is in our first manga
		hasTag := false
		for _, tag2 := range manga.Attributes.Tags {
			if tag2.Id == tag1.Id {
				hasTag = true
				break
			}
		}

		// If we have the tag, then no need to check the other manga
		// If we don't have it, then the other manga shouldn't have it..
		if hasTag {
			continue
		}

		// Check if other does not have the tag
		for _, tag2 := range mangaOther.Attributes.Tags {
			if tag2.Id == tag1.Id {
				return true
			}
		}

	}

	// Small check for "promo" titles, don't match to promotional titles
	title := strings.ToLower((*manga.Attributes.Title)["en"])
	titleOther := strings.ToLower((*mangaOther.Attributes.Title)["en"])
	if !strings.Contains(title, "(promo)") && strings.Contains(titleOther, "(promo)") {
		return true
	}

	// Else this is a valid match we can use!
	return false

}
