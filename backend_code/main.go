package main

import (
	"culturehouse/database" // Импорт пакета с подключением к бд
	"culturehouse/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	defer database.DB.Close()

	fmt.Println("Успешно подключено к базе данных!")

	routes.SetupRoutes()
	log.Println("Сервер запущен на порту :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
