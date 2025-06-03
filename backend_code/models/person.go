package models

import (
	"database/sql"
)

type RoleType string

const (
	RoleTeacher   RoleType = "teacher"
	RoleChild     RoleType = "child"
	RoleEmployee  RoleType = "employee"
	RoleParent    RoleType = "parent"
	RoleApplicant RoleType = "applicant"
)

type Person struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Lastname    *string  `json:"lastname,omitempty"`
	GroupID     *int     `json:"group_id,omitempty"`
	PhoneNumber string   `json:"phone_number"`
	Role        RoleType `json:"role"`
	IsActive    bool     `json:"is_active"`
	VacancyID   *int     `json:"vacancy_id,omitempty"`
	Photo       *string  `json:"photo,omitempty"`
}

func GetTeacherByGroupIDWithPhoto(db *sql.DB, groupID int) (*Person, error) {
	query := `
		SELECT 
			p.id, 
			p.name, 
			p.surname, 
			p.lastname, 
			p.phone_number, 
			COALESCE(m.file_path, '')
		FROM person p
		LEFT JOIN media_with_person m ON m.related_table = 'person' AND m.related_id = p.id
		WHERE p.role = 'teacher' AND p.id_group = $1
		LIMIT 1
	`

	var teacher Person
	err := db.QueryRow(query, groupID).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Surname,
		&teacher.Lastname,
		&teacher.PhoneNumber,
		&teacher.Photo,
	)

	if err == sql.ErrNoRows {
		return nil, nil // преподаватель не найден
	}
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func GetAllContactPersons(db *sql.DB) ([]Person, error) {
	query := `
		SELECT DISTINCT ON (p.id)
			p.id,
			p.name,
			p.surname,
			p.lastname,
			p.id_group,
			p.phone_number,
			p.role,
			p.is_active,
			p.vacancy_id,
			m.file_path
		FROM view_employee p
		LEFT JOIN media_with_person m ON m.related_id = p.id
		LEFT JOIN vacancy v ON p.vacancy_id = v.id
		ORDER BY p.id, m.uploaded_at ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Surname,
			&p.Lastname,
			&p.GroupID,
			&p.PhoneNumber,
			&p.Role,
			&p.IsActive,
			&p.VacancyID,
			&p.Photo,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}
