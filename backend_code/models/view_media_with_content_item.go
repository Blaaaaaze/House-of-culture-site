package models

import "time"

type MediaWithContentItem struct {
	ID           int       `json:"id"`
	RelatedTable string    `json:"related_table"`
	RelatedID    int       `json:"related_id"`
	Type         MediaType `json:"type"` // ENUM: image, video, pdf
	FilePath     string    `json:"file_path"`
	Description  string    `json:"description"`
	UploadedAt   time.Time `json:"uploaded_at"`

	ContentName      string      `json:"content_name"`
	ShortDescription string      `json:"short_description"`
	ContentType      ContentType `json:"content_type"` // ENUM: news, event, festival
}
