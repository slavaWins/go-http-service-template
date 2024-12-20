package models

import (
	"gorm.io/gorm"
)


type Template struct {
	gorm.Model  // Встроенные поля: ID, CreatedAt, UpdatedAt, DeletedAt
    id int
    FirstName string
    LastName  string
}

// Метод для получения полного имени
func (u *Template) GetFullName() string {
       return u.FirstName + " " + u.LastName
}
