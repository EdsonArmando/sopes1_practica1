package controllers

import (
    "context"
    "encoding/json"
    "apiRest/configs"
    "apiRest/models"
    "apiRest/responses"
    "net/http"
    "time"
    "github.com/go-playground/validator/v10"   
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var operationCollection *mongo.Collection = configs.GetCollection(configs.DB, "operation")
var validate = validator.New()


func CreateOperation() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var operacion models.Operation
        defer cancel()

        //validate the request body
        if err := json.NewDecoder(r.Body).Decode(&operacion); err != nil {
            rw.WriteHeader(http.StatusBadRequest)
            response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&operacion); validationErr != nil {
            rw.WriteHeader(http.StatusBadRequest)
            response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        newOperation := models.Operation{
            Id:       primitive.NewObjectID(),
            Numero1:  operacion.Numero1,
            Numero2:   operacion.Numero2,
            Operacion:    operacion.Operacion,
            Resultado:    operacion.Resultado,
            Fecha:    operacion.Fecha,
        }
        result, err := operationCollection.InsertOne(ctx, newOperation)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        rw.WriteHeader(http.StatusCreated)
        response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
        json.NewEncoder(rw).Encode(response)
    }
}

func GetAllOperation() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.Operation
        defer cancel()

        results, err := operationCollection.Find(ctx, bson.M{})

        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleUser models.Operation
            if err = results.Decode(&singleUser); err != nil {
                rw.WriteHeader(http.StatusInternalServerError)
                response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
                json.NewEncoder(rw).Encode(response)
            }

            users = append(users, singleUser)
        }

        rw.WriteHeader(http.StatusOK)
        response := responses.UserResponse{Status: 201, Message: "Funcion", Data: map[string]interface{}{"data": users}}
        json.NewEncoder(rw).Encode(response)
    }
}
