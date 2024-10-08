package controls

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AarizZafar/Nexus_verify.git/model"
)

const DBname = "NexusBiometrics"               // DB Name 
const ColName = "Devices"                      // Collection name

var Client *mongo.Client
var Collection  *mongo.Collection 

func init() {
	fmt.Println("Establishing connection with Mongo DB ........")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")      // creating an instance of connection to the MongoDB server 
	Client, err := mongo.Connect(context.TODO(), clientOptions)                // connecting to the MongoDB server using client options 
	handleErr(err)
	fmt.Println("Connection established with the MongoDB server........")
	fmt.Printf("************************************************\n")

	fmt.Println("Checking the connection")
	err = Client.Ping(context.TODO(), nil)
	handleErr(err)

	fmt.Println("Successfull connection to DB")
	fmt.Printf("************************************************\n")

	fmt.Printf("Getting DB %s and the collection %s",DBname, ColName)
	Collection = getCollection(Client)
	fmt.Printf("Success in accessing the Data Base %s", DBname)
	fmt.Printf("\n************************************************\n")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getCollection(client *mongo.Client) *mongo.Collection {
	Collection := client.Database(DBname).Collection(ColName)

	/* RunCommand - executes a command on MongoDB {'ping' : 1} used to check the connection 
	   bson.D{{Key: "ping", Value: 1}} is a BSON (Binary JSON) document that MongoDB interprets as a “ping” request.
	*/
	err := Collection.Database().RunCommand(context.TODO(), bson.D{{Key:"ping", Value:1}}).Err()   // checks if the mongodb server is rechable by sending a simple "ping" command
	handleErr(err)

	if err != nil {
		// if the db does not exist the db will be CREATED automatically when we inster data into the DB
		fmt.Printf("the DataBase %s will be created when we insert data in the collection %s \n", DBname, ColName)
	}
	return Collection
}

// --------------------------------------

func VerifyBiometics(c *gin.Context) {
	fmt.Println("Currently the user is beeing added directly but later on the verification process will be added")
	var sysBioMetrics model.SysBioMetrix

	/* when sending the struct from the local machine 
	   BindJSON - Reads the JSON payload from the request 
	   Mapping the JSON field to the Go struct the we have defind in the Azure vm models 
	*/
	if err := c.BindJSON(&sysBioMetrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existData, err := checkBioMetrics(sysBioMetrics)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error checking data"})
		return
	}

	// the device biometrics is in the MongoDB, Biometrics
	if existData != nil {
		c.JSON(http.StatusOK, gin.H{"status":"Data verify in DB", "data" : existData})
	} else {
		if err := insertBioMetrics(sysBioMetrics); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failer to insert data"})
			return 
		}
		c.JSON(http.StatusOK, gin.H{"status" : "Data inserted"})
	}
}

func checkBioMetrics(sysBioMetrix model.SysBioMetrix) (*model.SysBioMetrix, error) {
	var result model.SysBioMetrix
	filter := bson.M{
		"MAC":               sysBioMetrix.MAC,
		"CPUSerial":         sysBioMetrix.CPUSerial,
		"HardDriveSerial":   sysBioMetrix.HardDriveSerial,
		"MotherBoardSerial": sysBioMetrix.MotherBoardSerial,
		"BIOSSerial":        sysBioMetrix.BIOSSerial,
		"SSDSerial":         sysBioMetrix.SSDSerial,
		"TPMChipID":         sysBioMetrix.TPMChipID,
		"RAMSerial":         sysBioMetrix.RAMSerial,
		"GPUSerial":         sysBioMetrix.GPUSerial,
		"NICID":             sysBioMetrix.NICID,
	}

	// query the database with the filter 
	err := Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil    // the document was not founf
		}
		return nil, err        // we have faced some other error
	}
	return &result, nil        // the document was found 
}

func insertBioMetrics(sysBioMetrix model.SysBioMetrix) error {
	fmt.Println("Inserting Device Bio Metrics")
	fmt.Println("bio received : ", sysBioMetrix)
	_, err := Collection.InsertOne(context.TODO(), sysBioMetrix)
	return err
}

// ------------------------------


