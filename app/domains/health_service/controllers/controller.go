package controllers

import (
    "demo/pkg/contracts"
    "net/http"
)


func HealthGetHandler(r *http.Request) contracts.Response[string] {


    response := contracts.Response[string]{Value: "Ok"}

   return response;
}
