package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Operation struct {
    Id             primitive.ObjectID 
    Numero1       string          
    Numero2       string            
    Operacion     string            
    Resultado     string             
    Fecha         string             
}