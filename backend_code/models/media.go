package models

import (
	"database/sql"
	"fmt"
	"time"
)

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaVideo MediaType = "video"
	MediaPDF   MediaType = "pdf"
)

type Media struct {
	ID           int       `json:"id"`
	RelatedTable *string   `json:"related_table"`
	RelatedID    *int      `json:"related_id"`
	Type         MediaType `json:"type"`
	FilePath     string    `json:"file_path"`
	Description  string    `json:"description"`
	UploadedAt   time.Time `json:"uploaded_at"`
}

func GetAllMedia(db *sql.DB) ([]Media, error) {
	query := `
		SELECT *
		FROM media
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var media []Media
	for rows.Next() {
		var m Media
		if err := rows.Scan(&m.ID, &m.RelatedTable, &m.RelatedID, &m.Type, &m.FilePath, &m.Description, &m.UploadedAt); err != nil {
			return nil, err
		}
		media = append(media, m)
	}

	return media, nil
}

func GetMediaByGroupID(db *sql.DB, groupID int) ([]MediaWithGroup, error) {
	query := `
		SELECT id, related_table, related_id, type, file_path, description,
		       uploaded_at, group_name, short_description
		FROM media_with_group
		WHERE related_id = $1
	`

	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaList []MediaWithGroup
	for rows.Next() {
		var m MediaWithGroup
		if err := rows.Scan(
			&m.ID,
			&m.RelatedTable,
			&m.RelatedID,
			&m.Type,
			&m.FilePath,
			&m.Description,
			&m.UploadedAt,
			&m.GroupName,
			&m.ShortDescription,
		); err != nil {
			return nil, err
		}
		mediaList = append(mediaList, m)
	}

	return mediaList, nil
}

func GetMediaForContentItem(db *sql.DB, contentID int) ([]MediaWithContentItem, error) {
	query := `
		SELECT id, related_table, related_id, type, file_path, description, uploaded_at, content_name, short_description, content_type
		FROM media_with_content_item
		WHERE related_id = $1
	`
	rows, err := db.Query(query, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var media []MediaWithContentItem
	for rows.Next() {
		var m MediaWithContentItem
		err := rows.Scan(
			&m.ID, &m.RelatedTable, &m.RelatedID, &m.Type, &m.FilePath, &m.Description,
			&m.UploadedAt, &m.ContentName, &m.ShortDescription, &m.ContentType,
		)
		if err != nil {
			return nil, err
		}
		media = append(media, m)
	}
	return media, nil
}

func GetMediaByContentID(db *sql.DB, contentID int) ([]MediaWithContentItem, error) {
	query := `
		SELECT id, related_table, related_id, type, file_path, description, 
		       uploaded_at, content_name, short_description, content_type
		FROM media_with_content_item
		WHERE related_id = $1
	`

	rows, err := db.Query(query, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var media []MediaWithContentItem
	for rows.Next() {
		var m MediaWithContentItem
		err := rows.Scan(
			&m.ID, &m.RelatedTable, &m.RelatedID, &m.Type,
			&m.FilePath, &m.Description, &m.UploadedAt,
			&m.ContentName, &m.ShortDescription, &m.ContentType,
		)
		if err != nil {
			return nil, err
		}
		media = append(media, m)
	}
	return media, nil
}

func GetPDFMediaByFolder(db *sql.DB, folder string, deep bool) ([]Media, error) {
	var query string
	var args []interface{}

	if deep {
		// Файлы с подпапками
		query = `
			SELECT id, related_table, related_id, type, file_path, description, uploaded_at
			FROM media
			WHERE type = $1 AND file_path LIKE $2
			ORDER BY uploaded_at DESC
		`
		args = []interface{}{MediaPDF, folder + "%"}
	} else {
		// Только файлы внутри папки, без подпапок
		query = `
			SELECT id, related_table, related_id, type, file_path, description, uploaded_at
			FROM media
			WHERE type = $1 AND file_path ~ $2
			ORDER BY uploaded_at DESC
		`
		regex := fmt.Sprintf("^%s[^/]+\\.pdf$", folder)
		args = []interface{}{MediaPDF, regex}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Media
	for rows.Next() {
		var m Media
		err := rows.Scan(&m.ID, &m.RelatedTable, &m.RelatedID, &m.Type, &m.FilePath, &m.Description, &m.UploadedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}
