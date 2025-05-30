package models

import "time"

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaVideo MediaType = "video"
	MediaPDF   MediaType = "pdf"
)

type Media struct {
	ID           int       `json:"id"`
	RelatedTable string    `json:"related_table"`
	RelatedID    int       `json:"related_id"`
	Type         MediaType `json:"type"`
	FilePath     string    `json:"file_path"`
	Description  string    `json:"description"`
	UploadedAt   time.Time `json:"uploaded_at"`
}
