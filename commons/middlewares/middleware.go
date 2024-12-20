package middlewares

import (
    "demo/commons/contracts"
    "encoding/json"
    "net/http"
)

type HandlerFunc[T any] func(r *http.Request) contracts.Response[T]

func Wrapper[T any](handler HandlerFunc[T]) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        response := handler(r)

        if(response.Status==0){
         response.Status = 200;
        }

        w.WriteHeader( response.Status)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
