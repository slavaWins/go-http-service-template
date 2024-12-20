package app

import (
    "demo/domains/template_service"
    "fmt"
)



func Routes() {

    fmt.Println("Routes init") // Вывод: Иван Иванов

    template_service.Routes()
}
