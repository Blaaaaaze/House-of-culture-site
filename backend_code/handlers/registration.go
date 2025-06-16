package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"encoding/json"
	"net/http"
)

func RegisterChildHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req models.Registration
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка чтения JSON", http.StatusBadRequest)
		return
	}

	err := models.RegisterChildWithParent(database.DB, req)
	if err != nil {
		http.Error(w, "Ошибка при сохранении: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ребёнок и родитель успешно зарегистрированы"))
}
