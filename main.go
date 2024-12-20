package main

import (
    "demo/app/controllers"
    "demo/commons/middlewares"
    "net/http"
    "fmt"
)

// Response структура для JSON-ответов
type Response struct {
    Message string `json:"message"`
}


func main() {

    fmt.Println("Routes init") // Вывод: Иван Иванов

   // http.HandleFunc("/", controllers.IndexHandler)
    http.HandleFunc("/", middlewares.Wrapper(controllers.Index))
    http.HandleFunc("/user", middlewares.Wrapper(controllers.IndexHandlerResult))


    fmt.Println("Started") // Вывод: Иван Иванов
    http.ListenAndServe(":8080", nil)

}
