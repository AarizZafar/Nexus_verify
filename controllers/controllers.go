package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AarizZafar/Nexus_verify.git/admins_auth"
	"github.com/AarizZafar/Nexus_verify.git/model"
)

var adminsCred = admins_auth.Admins_creds
var adminVerifyClient *mongo.Client

func init() {
	adminVerifyClient = GetMongoSession()
}

func GetMongoSession() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // creating an instance of connection to the MongoDB server
	Client, err := mongo.Connect(context.TODO(), clientOptions)             // connecting to the MongoDB server using client options
	handleErr(err)

	fmt.Println("\033[97;46m             CHECKING CONNECTION WITH MONGO DB                    \033[0m")
	err = Client.Ping(context.TODO(), nil)
	handleErr(err)
	fmt.Println("\033[97;42m               SUCCESSFUL CONNECTION TO DB                 ✔      \033[0m")

	return Client
}

func AdminLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login_page_admin.html", nil)
}

func AdminAuthentication(ctx *gin.Context) {
	adminName := ctx.PostForm("username")
	adminPass := ctx.PostForm("password")
	// network     := ctx.PostForm("network")
	// testNetName := ctx.PostForm("testnet_name")
	// testNetPass := ctx.PostForm("testnet_pass")

	passWd, exist := adminsCred[adminName]

	if exist && adminPass == passWd {
		dbName := "ZU212"     // network db name
		colName := "testnet1" // network db collections name

		// adminVerifyClient = GetMongoSession()
		data, err := GetdatafromDB(dbName, colName)
		if err != nil {
			handleErr(err)
			ctx.HTML(http.StatusInternalServerError, "failure.html", nil)
		}

		ctx.HTML(http.StatusOK, "success_login_admin.html", gin.H{
			"data": data,
		})
	} else {
		ctx.HTML(http.StatusUnauthorized, "failure.html", nil)
	}
}

func GetdatafromDB(dbName string, colName string) ([]bson.M, error) {
	if adminVerifyClient == nil {
		return nil, fmt.Errorf("MongoDB client is not initialized")
	}

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

/*
   In a GIN HTTP handler we dont return a value to the caller insted we send it back to the client
   The below function is an http handler function, it interacts with the client via the http response,
   not by returning a string slice

   main action is sending response directly to the client the response contains the SSID encoded in JSON
*/

// this function is returning all the ssid that have been registered

func GetSSIDS(c *gin.Context) {
	if adminVerifyClient == nil {
		fmt.Println("\033[41m     THE MONGO CLIENT IS NOT CONNECTED     \033[0m")
		return
	}

	NetCol := adminVerifyClient.Database("Networks").Collection("NetbioMetrix")
	cursor, err := NetCol.Find(context.TODO(), bson.M{})
	handleErr(err)

	defer cursor.Close(context.TODO())

	var SSIDS []string
	for cursor.Next(context.TODO()) {
		var record bson.M
		err := cursor.Decode(&record)
		handleErr(err)

		if ssid, ok := record["ssid"].(string); ok {
			SSIDS = append(SSIDS, ssid)
		}
	}
	c.JSON(http.StatusOK, SSIDS)
}

func GetAdminCred(c *gin.Context) {
	fmt.Println("----")
	if adminVerifyClient == nil {
		fmt.Println("\033[41m     THE MONGO CLIENT IS NOT CONNECTED     \033[0m")
		return
	}

	AdminCreds := admins_auth.Admins_creds
	if len(adminsCred) == 0 {
		fmt.Println("\033[41m     THE ADMINS HAVE NOT BEEN DEFINED     \033[0m")
		c.JSON(http.StatusNotFound, gin.H{"message ": " NO ADMINS DEFINED "})
		return
	}

	c.JSON(http.StatusOK, AdminCreds)
}

// The below function returns all the testnets that are there in the ssid / network
func GetTestNetsfromSSID(c *gin.Context) {
	if adminVerifyClient == nil {
		fmt.Println("\033[41m     THE MONGO CLIENT IS NOT CONNECTED     \033[0m")
		return
	}

	ssid := c.Query("ssid")
	if ssid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "SSID (database name) is required"})
		return
	}

	db := adminVerifyClient.Database(ssid)

	testNetCols, err := db.ListCollectionNames(context.TODO(), map[string]interface{}{})
	if err != nil {
		fmt.Println("\033[41m     ERROR FETCHING COLLECTIONS : ", err, "     \033[0m")
		return
	}

	if testNetCols == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve collections"})
		return
	}

	fmt.Println(testNetCols)
	c.JSON(http.StatusOK, gin.H{"collection": testNetCols})
}

// the below function is returning all biometrix registered in the testnet
func GetBioMetrixfromTestNet(c *gin.Context) {
	if adminVerifyClient == nil {
		fmt.Println("\033[41m     THE MONGO CLIENT IS NOT CONNECTED     \033[0m")
		return
	}
	ssid := c.Query("ssid")
	testNet := c.Query("testNet")

	if ssid == "" || testNet == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both ssid and testNet are required"})
		return
	}

	testNetCol := adminVerifyClient.Database(ssid).Collection(testNet)

	cursor, err := testNetCol.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("\033[41m     ERROR FETCHING DATA : ", err, "     \033[0m")
		return
	}

	var BioMetrix []bson.M

	if err := cursor.All(context.TODO(), &BioMetrix); err != nil {
		fmt.Println("\033[41m     ERROR DECODING DATA : ", err, "     \033[0m")
		return
	}

	if BioMetrix == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve data from database"})
	}
	c.JSON(http.StatusOK, gin.H{"records": BioMetrix})
}

func CrtTestNetInSSID(c *gin.Context) {
	var TestNetData struct {
		Ssid       string             `json:"ssid"`
		NewTestNet string             `json:"testNet"`
		SysBioMetx model.SysBioMetrix `json:"sysbioMetx"`
	}

	if err := c.ShouldBindJSON(&TestNetData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error unmarshaling JSON: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})

	colCrt := false
	bioMetxInserted := false

	ssid := TestNetData.Ssid
	newTestNet := TestNetData.NewTestNet
	sysBioMetx := TestNetData.SysBioMetx

	db := adminVerifyClient.Database(ssid)
	col := db.Collection(newTestNet)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := col.InsertOne(ctx, sysBioMetx)
	if err != nil {
		fmt.Println("Error inserting system biometrics:", err)
		return
	}

	colCrt = true
	bioMetxInserted = true

	if colCrt && bioMetxInserted {
		fmt.Printf("\033[42mNEW TESTNET '%s' CREATED IN SSID '%s' ✔   \033[0m\n", TestNetData.NewTestNet, TestNetData.Ssid)
	} else {
		fmt.Printf("\033[41m%-10s  ERROR CREATING TESTNET '%s' IN SSID '%s'  %-10s\033[0m\n", "", TestNetData.NewTestNet, TestNetData.Ssid, "")
	}
}

func handleErr(err error) {
	if err != nil {
		log.Fatal("\033[97;41m", err, "\033[0m")
	}
}
