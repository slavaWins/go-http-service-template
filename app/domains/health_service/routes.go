package health_service

import (
    "net/http"
    "demo/pkg/middlewares"
    "demo/app/domains/health_service/controllers"
)


func Routes(){


     http.HandleFunc("/health", middlewares.Wrapper(controllers.HealthGetHandler))
}
