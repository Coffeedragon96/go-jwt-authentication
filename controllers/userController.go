package controllers

import (
	"github.com/coffeedragon96/go-jwt-authentication/database"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func SignUp() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() {

}
