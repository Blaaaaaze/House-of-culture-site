package models

type Vacancy struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	LongDescription string `json:"long_description"`
	IsActive        bool   `json:"is_active"`
}
