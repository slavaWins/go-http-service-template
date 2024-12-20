package models



type User struct {
    id int
    FirstName string
    LastName  string
}

// Метод для получения полного имени
func (u *User) GetFullName() string {
       return u.FirstName + " " + u.LastName
}
