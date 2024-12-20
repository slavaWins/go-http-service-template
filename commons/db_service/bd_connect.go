package db_service

import (
    "demo/commons/helpers"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
	"os"
)

// Глобальная переменная для хранения объекта DB
var DB *gorm.DB

// Инициализация базы данных
func init() {

    urlConnection := "{DB_USERNAME}:{DB_PASSWORD}@tcp({DB_HOST}:{DB_PORT})/{DB_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local";

   urlConnection = strings.Replace(urlConnection, "{DB_USERNAME}", helpers.Env("DB_USERNAME"), -1)
   urlConnection = strings.Replace(urlConnection, "{DB_PASSWORD}", os.Getenv("DB_PASSWORD")  , -1)
   urlConnection = strings.Replace(urlConnection, "{DB_HOST}", helpers.Env("DB_HOST"), -1)
   urlConnection = strings.Replace(urlConnection, "{DB_PORT}", helpers.Env("DB_PORT"), -1)
   urlConnection = strings.Replace(urlConnection, "{DB_DATABASE}", helpers.Env("DB_DATABASE"), -1)
  // fmt.Println(urlConnection);


	// Подключение к базе данных
	db, err := gorm.Open(mysql.Open(urlConnection), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Проверка подключения
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Не удалось получить объект DB: %v", err)
	}
	//defer sqlDB.Close()

	// Ping для проверки соединения
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Ошибка пинга базы данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")

	DB = db
}

// Connect - функция для получения объекта DB
func Connect() *gorm.DB {
	return DB
}
