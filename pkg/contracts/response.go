package contracts

import (
  //  "encoding/json"
  //  "net/http"
)

// Response структура для JSON-ответов
type Response[T any] struct {
    Value T
    Error *string `json:"Error,omitempty"`
    Status int
}

func ShortError(msg string ) Response[string]{

    error1 := Response[string]{Error: &msg , Status: 500}

    return error1
}

func ShortErrorGeneric [T any](msg string) Response[T] {
    return Response[T]{Error: &msg, Status: 500}
}
