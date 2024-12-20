package template_service

import (
    "net/http"
    "demo/pkg/middlewares"
    "demo/pkg/db_service"
    "fmt"
    "demo/app/domains/template_service/models"
    "demo/app/domains/template_service/controllers"
)


func Routes(){

	 fmt.Println("[Migrate] TemplateService migration run ")
     db :=  db_service.Connect();
     db.AutoMigrate(&models.Template{})

     http.HandleFunc("/api/template-service/v1/template", middlewares.Wrapper(controllers.TemplateGetHandler))
}
