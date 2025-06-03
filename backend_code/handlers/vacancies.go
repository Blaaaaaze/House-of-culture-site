package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"encoding/json"
	"net/http"
)

func GetActiveVacanciesHandler(w http.ResponseWriter, r *http.Request) {
	vacancies, err := models.GetActiveVacancy(database.DB)
	if err != nil {
		http.Error(w, "Ошибка при получении вакансий", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancies)
}
