package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetGroupsHandler(w http.ResponseWriter, r *http.Request) {
	groups, err := models.GetActiveGroups(database.DB)
	if err != nil {
		http.Error(w, "Ошибка при получении групп", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func GetGroupByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/groups/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	group, err := models.GetGroupByID(database.DB, id)
	if err != nil {
		http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func GetGroupWithMediaHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/groups/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID группы", http.StatusBadRequest)
		return
	}

	// Получаем группу
	group, err := models.GetGroupByID(database.DB, id)
	if err != nil {
		http.Error(w, "Ошибка при получении группы", http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.NotFound(w, r)
		return
	}

	// Получаем медиа, связанные с этой группой
	media, err := models.GetMediaByGroupID(database.DB, id)
	if err != nil {
		http.Error(w, "Ошибка при получении медиа", http.StatusInternalServerError)
		return
	}

	// Получаем преподавателя с фото (из view)
	teacher, err := models.GetTeacherByGroupIDWithPhoto(database.DB, id)
	if err != nil {
		http.Error(w, "Ошибка при получении преподавателя", http.StatusInternalServerError)
		return
	}

	// Составляем финальный JSON
	response := map[string]interface{}{
		"group":   group,
		"media":   media,
		"teacher": teacher, // может быть nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
