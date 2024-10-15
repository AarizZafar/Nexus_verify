package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AdminLoginPage(ctx *gin.Context) {
	fmt.Println("----------------Admin Login Page-------------------")
	ctx.HTML(http.StatusOK, "login_page_admin.html", nil)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetMongoSession() *mongo.Client {
	fmt.Println("Establishing connection with Mongo DB ........")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")      // creating an instance of connection to the MongoDB server 
	Client, err := mongo.Connect(context.TODO(), clientOptions)                // connecting to the MongoDB server using client options 
	handleErr(err)

	fmt.Println("Connection established with the MongoDB server........")

	fmt.Println("Checking the connection")
	err = Client.Ping(context.TODO(), nil)
	handleErr(err)

	fmt.Println("Successfull connection to DB")

	return Client
}