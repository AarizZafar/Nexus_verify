package controlls

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBname = "NexusDB"               // DB Name 
const ColName = "Biometrics"           // Collection name
var Client *mongo.Client


func init() {
	fmt.Println("Establishing connection with Mongo DB ........")

	clientOptions := options.Client().ApplyURI("mongodb:localhost:27017")      // creating an instance of connection to the MongoDB server 
	Client, err := mongo.Connect(context.TODO(), clientOptions)                // connecting to the MongoDB server using client options 
	handleErr(err)

	fmt.Println("Connection established with the MongoDB server........")
	fmt.Println("Checking the connection")

	err = Client.Ping(context.TODO(), nil)
	handleErr(err)

	fmt.Println("Successfull connection to DB")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


