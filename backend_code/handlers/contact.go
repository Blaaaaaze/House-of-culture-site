package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"encoding/json"
	"net/http"
)

func GetContactPersonsHandler(w http.ResponseWriter, r *http.Request) {
	people, err := models.GetAllContactPersons(database.DB)
	if err != nil {
		http.Error(w, "Ошибка при получении контактов", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
