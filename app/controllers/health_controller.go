package controllers

import (
    "demo/commons/contracts"
    "net/http"
)


// Handler для корневого маршрута
func Index(r *http.Request) contracts.Response[string] {
   response := contracts.Response[string]{Value: "Ok"}
   return response;
}
