package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AarizZafar/Nexus_verify.git/admins_credentials"
)

var adminCred = admins_credentials.Admins_creds

func AdminLoginPage(ctx *gin.Context) {
	fmt.Println("----------------Admin Login Page-------------------")
	ctx.HTML(http.StatusOK, "login_page_admin.html", nil)
}

func AdminAuthentication(ctx *gin.Context) {
	fmt.Println("Admin Authentication ")

	adminName := ctx.PostForm("username")
	adminPass := ctx.PostForm("password")
	netWork := ctx.PostForm("network")
	testNetName := ctx.PostForm("testnet_name")
	testNetPass := ctx.PostForm("testnet_pass")

	fmt.Println("adminName : ", adminName, "\npassWord : ", adminPass, "\nnetWord : ", netWork, "\ntestNetName : ", testNetName, "\ntestNetPass : ", testNetPass)

	passWd, exist := adminCred[adminName]

	// changes have to be done here 
	if exist && adminPass == passWd {
		fmt.Println("The admin Credentials are correct")
		ctx.HTML(http.StatusOK, "success_login_admin.html", nil)
	} else {
		fmt.Println("The admin Credentials are wrong")
		ctx.HTML(http.StatusUnauthorized, "failure.html", nil)
	}

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