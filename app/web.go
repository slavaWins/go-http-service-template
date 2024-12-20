package app

import (
    "demo/app/domains/template_service"
    "demo/app/domains/health_service"
    "fmt"
)



func Routes() {

    fmt.Println("Routes init") // Вывод: Иван Иванов

    template_service.Routes()
    health_service.Routes()
}
