package story

// ChapterTitle is a simple type casting
// representing a ChapterTitle
type ChapterTitle string

// Story represents a whole story
type Story map[ChapterTitle]Chapter

// Chapter represents a section of a story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}
