package main

import (
    "demo/app"
    "demo/pkg/db_service"
    "demo/pkg/helpers"
    "net/http"
    "fmt"
    "github.com/joho/godotenv"
)



func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }


    helpers.Init();

    db_service.Connect();


    app.Routes();



    fmt.Println("Started") // Вывод: Иван Иванов


    http.ListenAndServe(":8080", nil)

}
