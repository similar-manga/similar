package similar

type SimilarMatch struct {
	Id            string            `json:"id,omitempty"`
	Title         map[string]string `json:"title,omitempty"`
	ContentRating string            `json:"contentRating,omitempty"`
	Score         float32           `json:"score,omitempty"`
	Languages     []string          `json:"languages,omitempty"`
}

type SimilarManga struct {
	Id             string            `json:"id,omitempty"`
	Title          map[string]string `json:"title,omitempty"`
	ContentRating  string            `json:"contentRating,omitempty"`
	SimilarMatches []SimilarMatch    `json:"matches,omitempty"`
	UpdatedAt      string            `json:"updatedAt,omitempty"`
}

type ChapterInformation struct {
	Id          string         `json:"id,omitempty"`
	NumChapters int            `json:"num_chapters,omitempty"`
	Languages   []string       `json:"languages,omitempty"`
	Groups      []ChapterGroup `json:"groups,omitempty"`
}

type ChapterGroup struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
