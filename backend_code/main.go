package main

import (
	"culturehouse/database" // Импорт пакета с подключением к бд
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Подключаемся к базе данных
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer db.Close()

	// Проверка соединения
	fmt.Println("Успешно подключено к базе данных!")

	// Запуск HTTP-сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// Слушаем на порту 80
	http.ListenAndServe(":80", nil)
}
