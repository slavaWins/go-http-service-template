package main

import (
    "demo/app"
    "demo/pkg/db_service"
   // "demo/pkg/controllers"
    "demo/pkg/helpers"
    "net/http"
    "fmt"
    "github.com/joho/godotenv"
   // "demo/pkg/middlewares"
    "github.com/gorilla/mux"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Ошибка 404: pidor Страница не найдена")
}

func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }


    helpers.Init();

    db_service.Connect();

    routes := mux.NewRouter()

    app.Routes(routes);

    http.Handle("/", routes)
    //http.HandleFunc("/",  middlewares.Wrapper(controllers.NotFoundHandler)  )

    fmt.Println("Started") // Вывод: Иван Иванов


    http.ListenAndServe(":8080", nil)

}
