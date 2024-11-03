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

func GetSSIDAdminLogin() []string {
	fmt.Print("\033[46m                 GET SSIDS ADMIN VERIFY                       \033[0m")
	admVrfyClt()

	NetCol := adminVerifyClient.Database("Networks").Collection("NetbioMetrix")
	cursor, err := NetCol.Find(context.TODO(), bson.M{})
	handleErr(err)

	defer cursor.Close(context.TODO())

	var SSID []string
	for cursor.Next(context.TODO()) {
		var record bson.M
		err := cursor.Decode(&record)
		handleErr(err)

		if ssid, ok := record["ssid"].(string); ok {
			SSID = append(SSID, ssid)
		}
	}
	fmt.Print("\033[42m ✔  \033[0m\n")

	return SSID
}

func GetTestNetsAdminLogin(ssid string) []string {
	fmt.Print("\033[46m                GET TEST NET ADMIN VERIFY                     \033[0m")
	admVrfyClt()

	testNetCols, err := adminVerifyClient.Database(ssid).ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		handleErr(err)
		return nil
	}

	fmt.Print("\033[42m ✔  \033[0m\n")
	return testNetCols
}

func GetBioMetxAdminlogin(ssid, testNetx string) ([]bson.M, error) {
	fmt.Print("\033[46m               GET BIO METX ADMIN VERIFY                      \033[0m")
	admVrfyClt()

	bioMetxcol := adminVerifyClient.Database(ssid).Collection(testNetx)
	cursor, err := bioMetxcol.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}

	defer cursor.Close(context.TODO())

	var records []bson.M
	for cursor.Next(context.TODO()) {
		var record bson.M
		if err := cursor.Decode(&record); err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		records = append(records, record)
	}

	fmt.Print("\033[42m ✔  \033[0m\n")
	return records, nil
}

func AdminAuthentication(ctx *gin.Context) {
	adminName 		:= ctx.PostForm("username")
	adminPass 		:= ctx.PostForm("password")
	network 		:= ctx.PostForm("network")
	testNetName 	:= ctx.PostForm("testnet_name")
	// testNetPass                := ctx.PostForm("testnet_pass")

	passWd, exist := adminsCred[adminName]

	var ssids = GetSSIDAdminLogin()
	var regSsid bool = false
	for _, ssid := range ssids {
		if network == ssid {
			regSsid = true
			break
		}
	}

	var testNets = GetTestNetsAdminLogin(network)
	var regTestNet bool = false
	for _, testNet := range testNets {
		if testNet == testNetName {
			regTestNet = true
			break
		}
	}

	var BioMetricsdata []bson.M
	var err error
	if regSsid && regTestNet {
		BioMetricsdata, err = GetBioMetxAdminlogin(network, testNetName)
		if err != nil {
			handleErr(err)
		}
	}

	if exist && adminPass == passWd {
		if regSsid && regTestNet {
			bioMetric := convertToBioMetric(BioMetricsdata)
			ctx.HTML(http.StatusOK, "success_login_admin.html", gin.H{
				"data": bioMetric,
			})
		} else {
			ctx.HTML(http.StatusUnauthorized, "failure.html", nil)
		}
	} else {
		ctx.HTML(http.StatusUnauthorized, "failure.html", nil)
	}
}

type BioMetric struct {
	ID                 string `json:"_id"`
	MAC                string `json:"mac"`
	SSID               string `json:"SSID"`
	SystemSerialNumber string `json:"systemserialnumber"`
	UUID               string `json:"uuid"`
}

func convertToBioMetric(data []bson.M) []BioMetric {
	var bioMetrics []BioMetric
	for _, item := range data {
			bioMetric := BioMetric{
				ID: 					fmt.Sprintf("%v", item["_id"]),
				MAC: 					fmt.Sprintf("%v", item["mac"]),
				SSID: 					fmt.Sprintf("%v", item["ssid"]),
				SystemSerialNumber: 	fmt.Sprintf("%v", item["systemserialnumber"]),
				UUID: 					fmt.Sprintf("%v", item["uuid"]),
		}
		bioMetrics = append(bioMetrics, bioMetric)
	}
	return bioMetrics
}

func GetdatafromDB(dbName string, colName string) ([]bson.M, error) {
	admVrfyClt()

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
	fmt.Print("\033[46m                           GET SSIDS                          \033[0m")
	admVrfyClt()

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
	fmt.Print("\033[42m ✔  \033[0m\n")
	c.JSON(http.StatusOK, SSIDS)
}

func GetAdminCred(c *gin.Context) {
	fmt.Print("\033[46m                       GET ADMIN CRED                         \033[0m")
	admVrfyClt()

	AdminCreds := admins_auth.Admins_creds
	if len(adminsCred) == 0 {
		fmt.Println("\033[41m     THE ADMINS HAVE NOT BEEN DEFINED     \033[0m")
		c.JSON(http.StatusNotFound, gin.H{"message ": " NO ADMINS DEFINED "})
		return
	}

	fmt.Print("\033[42m ✔  \033[0m\n")
	c.JSON(http.StatusOK, AdminCreds)
}

// The below function returns all the testnets that are there in the ssid / network
func GetTestNetsfromSSID(c *gin.Context) {
	fmt.Print("\033[46m                    GET TEST NETS FROM SSID                   \033[0m")
	admVrfyClt()

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

	fmt.Print("\033[42m ✔  \033[0m\n")
	c.JSON(http.StatusOK, gin.H{"collection": testNetCols})
}

// the below function is returning all biometrix registered in the testnet
func GetBioMetrixfromTestNet(c *gin.Context) {
	fmt.Print("\033[46m                GET BIOMETRIX FROM TEST NET               \033[0m")
	admVrfyClt()

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

	fmt.Print("\033[42m ✔  \033[0m\n")
	c.JSON(http.StatusOK, gin.H{"records": BioMetrix})
}

func CrtTestNetInSSID(c *gin.Context) {
	fmt.Print("\033[46m                      CRT TEST NET IN SSID                        \033[0m")

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

	db := adminVerifyClient.Database(TestNetData.Ssid)
	col := db.Collection(TestNetData.NewTestNet)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, TestNetData.SysBioMetx)
	if err != nil {
		fmt.Println("Error inserting system biometrics:", err)
		fmt.Printf("\033[41m%-10s  ERROR CREATING TESTNET '%s' IN SSID '%s'  %-10s\033[0m\n", "", TestNetData.NewTestNet, TestNetData.Ssid, "")
		return
	}

	fmt.Printf("\033[42m NEW TESTNET '%s' CREATED IN SSID '%s' ✔   \033[0m\n", TestNetData.NewTestNet, TestNetData.Ssid)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal("\033[97;41m", err, "\033[0m")
	}
}

func admVrfyClt() {
	if adminVerifyClient == nil {
		fmt.Println("\033[41m     THE MONGO CLIENT IS NOT CONNECTED     \033[0m")
		return
	}
}
