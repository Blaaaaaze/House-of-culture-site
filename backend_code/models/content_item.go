package models

type ContentType string

const (
	ContentNews     ContentType = "news"
	ContentEvent    ContentType = "event"
	ContentFestival ContentType = "festival"
)

type ContentItem struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	ShortDescription string      `json:"short_description"`
	FullDescription  string      `json:"full_description"`
	Type             ContentType `json:"type"`
	IsActive         bool        `json:"is_active"`
}
