package controllers

import (
    "demo/commons/contracts"
    "demo/commons/helpers"
    "net/http"
)


// Handler для корневого маршрута
func Index(r *http.Request) contracts.Response[string] {
   response := contracts.Response[string]{Value:  helpers.Env("APP_NAME") }
   return response;
}
