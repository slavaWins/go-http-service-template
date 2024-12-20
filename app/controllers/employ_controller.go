package controllers

import (
    "demo/commons/contracts"
    "net/http"
    "demo/app/models"
    "demo/commons/db_service"
)


func EmployGetHandler(r *http.Request) contracts.Response[models.Employ] {


   newProduct := models.Employ{
    	FirstName: "Tom",
    	LastName: "Braidi",
    }

    result := db_service.Connect().Create(&newProduct)
      if result.Error != nil {
            return contracts.ShortErrorGeneric[models.Employ]("Error create data: ");
      }


    response := contracts.Response[models.Employ]{Value: newProduct}

   return response;
}
