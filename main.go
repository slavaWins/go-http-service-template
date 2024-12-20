package main

import (
    "demo/commons/db_service"
    "demo/commons/helpers"
    "demo/domains/template_service"
    "net/http"
    "fmt"
    "github.com/joho/godotenv"
)

// Response структура для JSON-ответов
type Response struct {
    Message string `json:"message"`
}


func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }


    helpers.Init();

    db_service.Connect();

    template_service.Routes()



    fmt.Println("Started") // Вывод: Иван Иванов


    http.ListenAndServe(":8080", nil)

}
