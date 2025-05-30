package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Импорт для работы с виртуальным окружением
	_ "github.com/lib/pq"      // Импорт драйвера PostgreSQL
)

var DB *sql.DB

// Функция для подключения к базе данных
func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// Получаем переменные из окружения. Это файл .env
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Строка подключения к базе данных
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Проверяем, что подключение прошло успешно
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
