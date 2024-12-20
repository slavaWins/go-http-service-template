package controllers

import (
    "demo/pkg/contracts"
    "net/http"
    "demo/pkg/db_service"
    "demo/app/domains/template_service/models"
    "strconv"
    "github.com/gorilla/mux"
)


func CreateTemplateHandler(r *http.Request) contracts.Response[*models.Template] {

   newProduct := models.Template{
    	FirstName: "Tom",
    	LastName: "Braidi",
    }

    result := db_service.Connect().Create(&newProduct)
      if result.Error != nil {
            return contracts.ShortErrorGeneric[*models.Template]("Error create data ");
      }


    response := contracts.Response[*models.Template]{Value: &newProduct}

   return response;
}




func GetTemplateHandler(r *http.Request) contracts.Response[*models.Template] {

   vars := mux.Vars(r)
      idStr := vars["id"]

      id, err := strconv.Atoi(idStr)
      if err != nil {
         return contracts.ShortErrorGeneric[*models.Template]("Invalid ID format ");
      }

      // Логика получения модели по id
      // Например, запрос к базе данных с использованием полученного id
      template := models.Template{}
      db := db_service.Connect()
      result := db.First(&template, id)

      if result.Error != nil {
         return contracts.ShortErrorGeneric[*models.Template]("Template not found");
      }


    response := contracts.Response[*models.Template]{Value: &template}

   return response;
}


