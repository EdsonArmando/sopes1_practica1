package routes

import (
    "apiRest/controllers" //add this
    "github.com/gorilla/mux"
)

func UserRoute(router *mux.Router)  {
   router.HandleFunc("/operation", controllers.CreateOperation()).Methods("POST")
   router.HandleFunc("/allOperations", controllers.GetAllOperation()).Methods("GET")
}
