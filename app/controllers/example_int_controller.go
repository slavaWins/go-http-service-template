package controllers

import (
    "demo/commons/contracts"
    "net/http"
)


// Handler для корневого маршрута
func ExampleIntHandler(r *http.Request) contracts.Response[int] {
  // val := 55;

   // response := contracts.Response[int]{Value: val}
   response := contracts.ShortErrorGeneric[int]("Problems");
   return response;
}
