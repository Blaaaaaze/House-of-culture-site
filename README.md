
# Сайт для городского дома культуры

## 📌 Описание проекта
Веб-приложение для городского дома культуры. В данный момент в процессе разработки. Readme будет своевременно обновляться

## 🚀 Технологии
- **Фронтенд:** HTML, CSS, SASS, JavaScript, Vue.js
- **Бэкенд:** Go, PostgreSQL
- **Соблюдение ГОСТ 52872-2019**

## 🛠 Развёртывание

### 1. Клонирование репозитория
```bash
git clone https://github.com/Blaaaaaze/house-of-culture.git
cd house-of-culture
```

## 🔧 Запуск бэкенда

### Перейдите в папку
   ```bash
   cd backend_code
  ```

### Установите зависимости
```bash
   go mod tidy
  ```

### Создайте .env файл
```bash
   cp .env.example .env
  ```

### Заполните своими данными
Например
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=app_db
```
Dump базы данных по адресу backend_code/database/culture_house_db_dump

