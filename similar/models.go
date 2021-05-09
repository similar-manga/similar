package similar

type SimilarMatch struct {
	Id string `json:"id,omitempty"`
	Title map[string]string `json:"title,omitempty"`
	AlternativeCover string `json:"alternativeCover,omitempty"`
	ContentRating string `json:"contentRating,omitempty"`
	Score float32 `json:"score,omitempty"`
}

type SimilarManga struct {
	Id string `json:"id,omitempty"`
	Title map[string]string `json:"title,omitempty"`
	AlternativeCover string `json:"alternativeCover,omitempty"`
	ContentRating string `json:"contentRating,omitempty"`
	SimilarMatches []SimilarMatch `json:"matches,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}




