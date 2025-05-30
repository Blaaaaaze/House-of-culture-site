package models

import "time"

type MediaWithGroup struct {
	ID           int       `json:"id"`
	RelatedTable string    `json:"related_table"`
	RelatedID    int       `json:"related_id"`
	Type         MediaType `json:"type"`
	FilePath     string    `json:"file_path"`
	Description  string    `json:"description"`
	UploadedAt   time.Time `json:"uploaded_at"`

	GroupName        string `json:"group_name"`
	ShortDescription string `json:"short_description"`
}
