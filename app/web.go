package app

import (
    "demo/app/domains/template_service"
   // "demo/app/domains/health_service"
    "fmt"
    "github.com/gorilla/mux"
)



func Routes(routes *mux.Router) {

    fmt.Println("Routes init") // Вывод: Иван Иванов

    template_service.Routes(routes)
   // health_service.Routes(routes)
}
