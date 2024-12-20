package health_service

import (
    "net/http"
    "demo/pkg/middlewares"
    "demo/app/domains/health_service/controllers"
    "github.com/gorilla/mux"
)


func Routes(routes *mux.Router){


     http.HandleFunc("/health", middlewares.Wrapper(controllers.HealthGetHandler))
}
