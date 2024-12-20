package controllers

import (
    "demo/pkg/contracts"
    "net/http"
)


func NotFoundHandler(r *http.Request) contracts.Response[*bool] {


   msg := "404 Error api method";

   return contracts.Response[*bool]{Status:404, Error: &msg};
}
