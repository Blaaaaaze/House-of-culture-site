package models

import (
	"database/sql"
)

type Group struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	FullDescription  string `json:"full_description"`
	IsActive         bool   `json:"is_active"`
	PreviewImage     string `json:"preview_image,omitempty"`
}

func GetActiveGroups(db *sql.DB) ([]Group, error) {
	query := `
		SELECT 
			g.id, g.name, g.short_description, g.full_description, g.is_active,
			(
				SELECT m.file_path
				FROM media_with_group m
				WHERE m.related_id = g.id AND m.type = 'image'
				ORDER BY m.uploaded_at ASC
				LIMIT 1
			) AS preview_image
		FROM "group" g
		WHERE g.is_active = true
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var g Group
		if err := rows.Scan(
			&g.ID, &g.Name, &g.ShortDescription, &g.FullDescription,
			&g.IsActive, &g.PreviewImage,
		); err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	return groups, nil
}

func GetGroupByID(db *sql.DB, id int) (*Group, error) {
	query := `
		SELECT id, name, short_description, full_description, is_active
		FROM "group"
		WHERE id = $1
	`

	var g Group
	err := db.QueryRow(query, id).Scan(&g.ID, &g.Name, &g.ShortDescription, &g.FullDescription, &g.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &g, nil
}
