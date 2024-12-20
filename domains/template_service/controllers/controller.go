package controllers

import (
    "demo/commons/contracts"
    "net/http"
    "demo/commons/db_service"
    "demo/domains/template_service/models"
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
