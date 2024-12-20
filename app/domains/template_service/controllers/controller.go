package controllers

import (
    "demo/pkg/contracts"
    "net/http"
    "demo/pkg/db_service"
    "demo/app/domains/template_service/models"
)


func TemplateGetHandler(r *http.Request) contracts.Response[models.Template] {



   newProduct := models.Template{
    	FirstName: "Tom",
    	LastName: "Braidi",
    }

    result := db_service.Connect().Create(&newProduct)
      if result.Error != nil {
            return contracts.ShortErrorGeneric[models.Template]("Error create data: ");
      }


    response := contracts.Response[models.Template]{Value: newProduct}

   return response;
}


