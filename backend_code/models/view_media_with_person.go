package models

import (
	"time"
)

type MediaWithPerson struct {
	ID           int       `json:"id"`
	RelatedTable string    `json:"related_table"`
	RelatedID    int       `json:"related_id"`
	Type         MediaType `json:"type"`
	FilePath     string    `json:"file_path"`
	Description  string    `json:"description"`
	UploadedAt   time.Time `json:"uploaded_at"`

	PersonName  string  `json:"person_name"`
	Surname     string  `json:"surname"`
	Lastname    *string `json:"lastname,omitempty"`
	PhoneNumber string  `json:"phone_number"`
}
