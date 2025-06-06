package models

import "database/sql"

type ContentType string

const (
	ContentNews     ContentType = "news"
	ContentEvent    ContentType = "event"
	ContentFestival ContentType = "festival"
)

type ContentItem struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	ShortDescription string      `json:"short_description"`
	FullDescription  string      `json:"full_description"`
	Type             ContentType `json:"type"`
	IsActive         bool        `json:"is_active"`
	PreviewImage     *string     `json:"preview_image,omitempty"`
}

func GetContentItemByID(db *sql.DB, id int) (*ContentItem, error) {
	query := `SELECT id, name, short_description, full_description, type, is_active FROM content_item WHERE id = $1`
	var c ContentItem
	err := db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.ShortDescription, &c.FullDescription, &c.Type, &c.IsActive)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetActiveNews(db *sql.DB) ([]ContentItem, error) {
	return getAllActiveContentByType(db, "news")
}

func GetActiveEvents(db *sql.DB) ([]ContentItem, error) {
	return getAllActiveContentByType(db, "event")
}

func GetActiveFestivals(db *sql.DB) ([]ContentItem, error) {
	return getAllActiveContentByType(db, "festival")
}

func getAllActiveContentByType(db *sql.DB, contentType string) ([]ContentItem, error) {
	query := `
		SELECT 
			ci.id, ci.name, ci.short_description, ci.full_description,
			ci.type, ci.is_active,
			(
				SELECT m.file_path
				FROM media_with_content_item m
				WHERE m.related_id = ci.id AND m.type = 'image'
				ORDER BY m.uploaded_at ASC
				LIMIT 1
			) AS preview_image
		FROM content_item ci
		WHERE ci.type = $1 AND ci.is_active = true
	`

	rows, err := db.Query(query, contentType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ContentItem
	for rows.Next() {
		var c ContentItem
		err := rows.Scan(
			&c.ID, &c.Name, &c.ShortDescription, &c.FullDescription,
			&c.Type, &c.IsActive, &c.PreviewImage,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, c)
	}
	return items, nil
}

func GetContentItemByTypeAndID(db *sql.DB, contentType string, id int) (*ContentItem, error) {
	query := `
		SELECT id, name, short_description, full_description, type, is_active
		FROM content_item
		WHERE id = $2 AND type = $1 AND is_active = true
	`

	var c ContentItem
	err := db.QueryRow(query, contentType, id).Scan(
		&c.ID, &c.Name, &c.ShortDescription, &c.FullDescription, &c.Type, &c.IsActive,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}
