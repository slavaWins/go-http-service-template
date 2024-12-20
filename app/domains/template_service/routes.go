package template_service

import (
    "demo/pkg/middlewares"
    "demo/pkg/db_service"
    "fmt"
    "demo/app/domains/template_service/models"
    "demo/app/domains/template_service/controllers"
    "github.com/gorilla/mux"
)


func Routes(routes *mux.Router){

	 fmt.Println("[Migrate] TemplateService migration run ")
     db :=  db_service.Connect();
     db.AutoMigrate(&models.Template{})

    // routes.HandleFunc("/v1/create/template", middlewares.Wrapper(controllers.CreateTemplateHandler))
     routes.HandleFunc("/v1/template/{id}", middlewares.Wrapper(controllers.GetTemplateHandler))
}
