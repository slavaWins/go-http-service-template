package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model  // Встроенные поля: ID, CreatedAt, UpdatedAt, DeletedAt
	Description string
	Price       float64
}

func main() {
	// Строка подключения к базе данных
	dsn := "root:@tcp(127.0.0.1:3306)/blank_tracker?charset=utf8mb4&parseTime=True&loc=Local"

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
	defer sqlDB.Close()

	// Ping для проверки соединения
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Ошибка пинга базы данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")


	err = db.AutoMigrate(&Product{})
    if err != nil {
    	log.Fatalf("Ошибка автомиграции: %v", err)
    }

    fmt.Println("Успешная миграция!")






    newProduct := Product{
    	Description: "Описание нового продукта",
    	Price:       99.99,
    }

    // Создание записи в базе данных
    result := db.Create(&newProduct)
    if result.Error != nil {
    	log.Fatalf("Ошибка при создании продукта: %v", result.Error)
    }

    fmt.Printf("Создан продукт с ID: %d\n", newProduct.ID)




    var products []Product

    // Получение всех продуктов
    resultList := db.Find(&products)
    if resultList.Error != nil {
    	log.Fatalf("Ошибка при получении продуктов: %v", resultList.Error)
    }

    fmt.Println("Список продуктов:")
    for _, p := range products {
    	fmt.Printf("ID: %d, Название: %s\n", p.ID, p.Description)
    }
}
