package models

import (
	"gorm.io/gorm"
)


type Employ struct {
	gorm.Model  // Встроенные поля: ID, CreatedAt, UpdatedAt, DeletedAt
    id int
    FirstName string
    LastName  string
}

// Метод для получения полного имени
func (u *Employ) GetFullName() string {
       return u.FirstName + " " + u.LastName
}
