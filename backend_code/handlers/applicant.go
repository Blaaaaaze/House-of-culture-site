package handlers

import (
	"culturehouse/database"
	"culturehouse/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func CreateApplicantHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Лимит в 10МБ
	if err != nil {
		http.Error(w, "Ошибка при обработке формы", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	lastname := r.FormValue("lastname")
	phone := r.FormValue("phone_number")
	vacancyID, err := strconv.Atoi(r.FormValue("vacancy_id"))
	if err != nil {
		http.Error(w, "Неверный ID вакансии", http.StatusBadRequest)
		return
	}

	personID, err := models.InsertPerson(database.DB, name, surname, lastname, phone, vacancyID)
	if err != nil {
		http.Error(w, "Ошибка записи в person", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("resume")
	if err != nil {
		http.Error(w, "Ошибка загрузки файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := fmt.Sprintf("uploads/docs/resumes/%d_%s", personID, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	desctiption := strconv.Itoa(personID) + "_Резюме"
	err = models.InsertMedia(database.DB, "person", personID, "pdf", filePath, desctiption)
	if err != nil {
		http.Error(w, "Ошибка записи в media", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Резюме успешно отправлено!")
}
