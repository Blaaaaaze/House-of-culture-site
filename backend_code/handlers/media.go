package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"encoding/json"
	"net/http"
)

func GetPDFMediaHandler(w http.ResponseWriter, r *http.Request) {
	folder := r.URL.Query().Get("folder")
	if folder == "" {
		folder = "/uploads/docs/" // Значение по умолчанию
	}

	deep := true
	if r.URL.Query().Get("deep") == "false" {
		deep = false
	}

	media, err := models.GetPDFMediaByFolder(database.DB, folder, deep)
	if err != nil {
		http.Error(w, "Ошибка при получении PDF-файлов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(media)
}
