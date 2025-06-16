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
	PhoneNumber *string  `json:"phone_number"`
	Role        RoleType `json:"role"`
	IsActive    bool     `json:"is_active"`
	VacancyID   *int     `json:"vacancy_id,omitempty"`
	Photo       *string  `json:"photo,omitempty"`
	Vacancy     *string  `json:"vacancy,omitempty"`
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
			p.role::text,
			p.is_active,
			p.vacancy_id,
			m.file_path,
			v.name AS vacancy_name
		FROM view_employee p
		LEFT JOIN media_with_person m ON m.related_id = p.id
		LEFT JOIN vacancy v ON p.vacancy_id = v.id
		WHERE v.id <= 7
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
			&p.Vacancy,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func InsertPerson(db *sql.DB, name, surname, lastname, phone string, vacancyID int) (int, error) {
	query := `
		INSERT INTO person (name, surname, lastname, phone_number, role, is_active, vacancy_id)
		VALUES ($1, $2, $3, $4, 'applicant', false, $5)
		RETURNING id`
	var id int
	err := db.QueryRow(query, name, surname, lastname, phone, vacancyID).Scan(&id)
	return id, err
}
