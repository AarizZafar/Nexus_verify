package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AarizZafar/Nexus_verify.git/admins_auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var adminsCred = admins_auth.Admins_creds
var adminVerifyClient *mongo.Client

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

func AdminLoginPage(ctx *gin.Context) {
	fmt.Println("----------------Admin Login Page-------------------")
	ctx.HTML(http.StatusOK, "login_page_admin.html", nil)
}

func AdminAuthentication(ctx *gin.Context) {
    fmt.Println("Admin Authentication function")

    adminName   := ctx.PostForm("username")
    adminPass   := ctx.PostForm("password")
    // network     := ctx.PostForm("network")
    // testNetName := ctx.PostForm("testnet_name")
    // testNetPass := ctx.PostForm("testnet_pass")

    passWd, exist := adminsCred[adminName]

    // checking if the admins credentials are correct 
    if exist && adminPass == passWd {
        fmt.Println("the credentials are correct")
        dbName  := "ZU212"       // network db name
        colName := "testnet1"    // network db collections name

        adminVerifyClient = GetMongoSession()
        data, err := GetdatafromDB(dbName, colName)
        if err != nil {
            fmt.Println("Error retriving data : ", err)
            ctx.HTML(http.StatusInternalServerError, "failure.html", nil)
        }

        ctx.HTML(http.StatusOK, "success_login_admin.html", gin.H{
            "data" : data,
        })
    } else {
        fmt.Println("The admin Credentials are wrong")
        ctx.HTML(http.StatusUnauthorized, "failure.html", nil)
    }
}

func GetdatafromDB(dbName string, colName string) ([]bson.M, error) {
    if adminVerifyClient == nil {
        return nil, fmt.Errorf("MongoDB client is not initialized")
    }

    fmt.Println(dbName, " : ", colName)
    collection := adminVerifyClient.Database(dbName).Collection(colName)
    cursor, err := collection.Find(context.TODO(), bson.M{})
    handleErr(err)

    defer cursor.Close(context.TODO())

    var records []bson.M
    for cursor.Next(context.TODO()) {
        var results bson.M
        err := cursor.Decode(&results)
        handleErr(err)

        records = append(records, results)
    }
    return records, nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

