package routes

import (
    "apiRest/controllers" //add this
    "github.com/gorilla/mux"
    "net/http"
)
func testApi(res http.ResponseWriter, req *http.Request){
    res.Write([]byte("Hello World"))
}
func UserRoute(router *mux.Router)  {
   router.HandleFunc("/operation", controllers.CreateOperation()).Methods("POST")
   router.HandleFunc("/allOperations", controllers.GetAllOperation()).Methods("GET")
   router.HandleFunc("/",testApi).Methods("GET");
}
