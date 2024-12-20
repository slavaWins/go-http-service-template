package main

import (
    "demo/app/controllers"
    "demo/commons/middlewares"
    "demo/commons/db_service"
    "demo/commons/migration"
    "demo/commons/helpers"
    "demo/domains/template_service"
    "net/http"
    "fmt"
)

// Response структура для JSON-ответов
type Response struct {
    Message string `json:"message"`
}


func main() {


    helpers.Init();

    db_service.Connect();

    migration.Migrate();

    fmt.Println(helpers.Env("APP_NAME")) // Вывод: Иван Иванов
    fmt.Println("Routes init") // Вывод: Иван Иванов

   // http.HandleFunc("/", controllers.IndexHandler)
   // http.HandleFunc("/", middlewares.Wrapper(controllers.Index))
    http.HandleFunc("/user", middlewares.Wrapper(controllers.EmployGetHandler))
    http.HandleFunc("/int", middlewares.Wrapper(controllers.ExampleIntHandler))

    template_service.Routes()


    fmt.Println("Started") // Вывод: Иван Иванов
    http.ListenAndServe(":8080", nil)

}
