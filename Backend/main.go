package main

import (
    "log"
    "apiRest/configs" //add this
    "apiRest/routes"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    headers := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT","DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"})
    //run database
    configs.ConnectDB()

    //routes
    routes.UserRoute(router) //add this

    log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers,methods,origins)(router)))
}