package models

type Group struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
	IsActive         bool   `json:"is_active"`
}
