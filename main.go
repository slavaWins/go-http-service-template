package main

import (
    "demo/app/controllers"
    "demo/commons/middlewares"
    "demo/commons/helpers"
    "net/http"
    "fmt"
)

// Response структура для JSON-ответов
type Response struct {
    Message string `json:"message"`
}


func main() {

    helpers.Init();


    fmt.Println(helpers.Env("APP_NAME")) // Вывод: Иван Иванов
    fmt.Println("Routes init") // Вывод: Иван Иванов

   // http.HandleFunc("/", controllers.IndexHandler)
   // http.HandleFunc("/", middlewares.Wrapper(controllers.Index))
    http.HandleFunc("/user", middlewares.Wrapper(controllers.IndexHandlerResult))


    fmt.Println("Started") // Вывод: Иван Иванов
    http.ListenAndServe(":8080", nil)

}
