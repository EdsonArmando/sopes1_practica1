package main

import (
    "log"
    "apiRest/configs" //add this
    "apiRest/routes"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    //run database
    configs.ConnectDB()

    //routes
    routes.UserRoute(router) //add this

    log.Fatal(http.ListenAndServe(":6000", router))
}