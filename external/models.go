package external

import "time"

type ResponseAniList struct {
	Data struct {
		Media struct {
			Coverimage struct {
				Extralarge string `json:"extraLarge"`
			} `json:"coverImage"`
		} `json:"Media"`
	} `json:"data"`
}

type ResponseAnimeList struct {
	RequestHash        string `json:"request_hash"`
	RequestCached      bool   `json:"request_cached"`
	RequestCacheExpiry int    `json:"request_cache_expiry"`
	Pictures           []struct {
		Large string `json:"large"`
		Small string `json:"small"`
	} `json:"pictures"`
}


type ResponseKitsuSearch struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			Createdat           time.Time `json:"createdAt"`
			Updatedat           time.Time `json:"updatedAt"`
			Slug                string    `json:"slug"`
			Synopsis            string    `json:"synopsis"`
			Description         string    `json:"description"`
			Coverimagetopoffset int       `json:"coverImageTopOffset"`
			Titles              struct {
				En   interface{} `json:"en"`
				EnJp string      `json:"en_jp"`
				EnUs string      `json:"en_us"`
				JaJp string      `json:"ja_jp"`
			} `json:"titles"`
			Canonicaltitle    string        `json:"canonicalTitle"`
			Abbreviatedtitles []interface{} `json:"abbreviatedTitles"`
			Averagerating     interface{}   `json:"averageRating"`
			Ratingfrequencies struct {
				Num2  string `json:"2"`
				Num3  string `json:"3"`
				Num4  string `json:"4"`
				Num5  string `json:"5"`
				Num6  string `json:"6"`
				Num7  string `json:"7"`
				Num8  string `json:"8"`
				Num9  string `json:"9"`
				Num10 string `json:"10"`
				Num11 string `json:"11"`
				Num12 string `json:"12"`
				Num13 string `json:"13"`
				Num14 string `json:"14"`
				Num15 string `json:"15"`
				Num16 string `json:"16"`
				Num17 string `json:"17"`
				Num18 string `json:"18"`
				Num19 string `json:"19"`
				Num20 string `json:"20"`
			} `json:"ratingFrequencies"`
			Usercount      int         `json:"userCount"`
			Favoritescount int         `json:"favoritesCount"`
			Startdate      string      `json:"startDate"`
			Enddate        string      `json:"endDate"`
			Nextrelease    interface{} `json:"nextRelease"`
			Popularityrank int         `json:"popularityRank"`
			Ratingrank     interface{} `json:"ratingRank"`
			Agerating      interface{} `json:"ageRating"`
			Ageratingguide interface{} `json:"ageRatingGuide"`
			Subtype        string      `json:"subtype"`
			Status         string      `json:"status"`
			Tba            interface{} `json:"tba"`
			Posterimage    struct {
				Tiny     string `json:"tiny"`
				Small    string `json:"small"`
				Medium   string `json:"medium"`
				Large    string `json:"large"`
				Original string `json:"original"`
				Meta     struct {
					Dimensions struct {
						Tiny struct {
							Width  interface{} `json:"width"`
							Height interface{} `json:"height"`
						} `json:"tiny"`
						Small struct {
							Width  interface{} `json:"width"`
							Height interface{} `json:"height"`
						} `json:"small"`
						Medium struct {
							Width  interface{} `json:"width"`
							Height interface{} `json:"height"`
						} `json:"medium"`
						Large struct {
							Width  interface{} `json:"width"`
							Height interface{} `json:"height"`
						} `json:"large"`
					} `json:"dimensions"`
				} `json:"meta"`
			} `json:"posterImage"`
			Coverimage    interface{} `json:"coverImage"`
			Chaptercount  int         `json:"chapterCount"`
			Volumecount   int         `json:"volumeCount"`
			Serialization interface{} `json:"serialization"`
			Mangatype     string      `json:"mangaType"`
		} `json:"attributes"`
		Relationships struct {
			Genres struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"genres"`
			Categories struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"categories"`
			Castings struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"castings"`
			Installments struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"installments"`
			Mappings struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mappings"`
			Reviews struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"reviews"`
			Mediarelationships struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mediaRelationships"`
			Characters struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"characters"`
			Staff struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"staff"`
			Productions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"productions"`
			Quotes struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"quotes"`
			Chapters struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"chapters"`
			Mangacharacters struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mangaCharacters"`
			Mangastaff struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"mangaStaff"`
		} `json:"relationships"`
	} `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}






