package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetNewsHandler(w http.ResponseWriter, r *http.Request) {
	news, err := models.GetActiveNews(database.DB)
	if err != nil {
		http.Error(w, "Ошибка при получении новостей", http.StatusInternalServerError)
		return
	}
	writeJSON(w, news)
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := models.GetActiveEvents(database.DB)
	if err != nil {
		http.Error(w, "Ошибка при получении событий", http.StatusInternalServerError)
		return
	}
	writeJSON(w, events)
}

// Получаем все активные фестивали
func GetFestivalsHandler(w http.ResponseWriter, r *http.Request) {
	fests, err := models.GetActiveFestivals(database.DB)
	if err != nil {
		http.Error(w, "Ошибка при получении фестивалей", http.StatusInternalServerError)
		return
	}
	writeJSON(w, fests)
}

// Получаем конкретную запись и все медиа к ней
func GetContentItemWithMediaHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 {
		http.Error(w, "Некорректный путь", http.StatusBadRequest)
		return
	}

	contentType := parts[1] // news, event, festival
	idStr := parts[2]

	// Проверка типа
	if contentType != "news" && contentType != "event" && contentType != "festival" {
		http.Error(w, "Неверный тип контента", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	// Получаем контент по типу и ID
	content, err := models.GetContentItemByTypeAndID(database.DB, contentType, id)
	if err != nil {
		http.Error(w, "Ошибка при получении контента", http.StatusInternalServerError)
		return
	}
	if content == nil {
		http.NotFound(w, r)
		return
	}

	// Получаем медиа
	media, err := models.GetMediaByContentID(database.DB, id)
	if err != nil {
		http.Error(w, "Ошибка при получении медиа", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"content": content,
		"media":   media,
	}

	writeJSON(w, response)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
