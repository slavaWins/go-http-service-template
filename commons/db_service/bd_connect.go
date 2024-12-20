package db_service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// Глобальная переменная для хранения объекта DB
var DB *gorm.DB

// Инициализация базы данных
func init() {
	// Строка подключения к базе данных
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/blank_tracker?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// Подключение к базе данных
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
