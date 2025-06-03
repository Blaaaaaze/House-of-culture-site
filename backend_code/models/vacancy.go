package models

import "database/sql"

type Vacancy struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	LongDescription string `json:"long_description"`
	IsActive        bool   `json:"is_active"`
}

func GetActiveVacancy(db *sql.DB) ([]Vacancy, error) {
	query := `
		SELECT id, name, long_description, is_active
		FROM vacancy
		WHERE is_active = true
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []Vacancy
	for rows.Next() {
		var v Vacancy
		if err := rows.Scan(&v.ID, &v.Name, &v.LongDescription, &v.IsActive); err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}

	return vacancies, nil
}
