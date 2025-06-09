package models

import (
	"database/sql"
)

type Registration struct {
	ChildName     string `json:"child_name"`
	ChildSurname  string `json:"child_surname"`
	ChildLastname string `json:"child_lastname,omitempty"`

	ParentName     string `json:"parent_name"`
	ParentSurname  string `json:"parent_surname"`
	ParentLastname string `json:"parent_lastname,omitempty"`
	ParentPhone    string `json:"parent_phone"`
}

func RegisterChildWithParent(db *sql.DB, req Registration) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Родитель в person
	var parentID int
	err = tx.QueryRow(`
		INSERT INTO person (name, surname, lastname, phone_number, role, is_active)
		VALUES ($1, $2, $3, $4, 'parent', true)
		RETURNING id
	`, req.ParentName, req.ParentSurname, req.ParentLastname, req.ParentPhone).Scan(&parentID)
	if err != nil {
		return err
	}

	// Ребёнок в person
	var childID int
	err = tx.QueryRow(`
		INSERT INTO person (name, surname, lastname, role, is_active)
		VALUES ($1, $2, $3, 'child', true)
		RETURNING id
	`, req.ChildName, req.ChildSurname, req.ChildLastname).Scan(&childID)
	if err != nil {
		return err
	}

	// Добавить в child_parent
	_, err = tx.Exec(`
		INSERT INTO child_parent (id_parent, id_child)
		VALUES ($1, $2)
	`, parentID, childID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
