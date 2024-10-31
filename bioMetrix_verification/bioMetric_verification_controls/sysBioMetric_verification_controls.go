package bioMetric_verification_controls

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/AarizZafar/Nexus_verify.git/controllers"
	"github.com/AarizZafar/Nexus_verify.git/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var SysBioNetDB = ""                                 // DB Name 
const SysBioNetCol = "testnet1"                      // Collection name

var SysBioClient *mongo.Client
var SysBioCol  *mongo.Collection 

func init() {
	SysBioClient = controllers.GetMongoSession()
}

func handleError(err error) {
	log.Fatal(err)
}

func dbColVerify(dbName string) {
	SysBioNetDB = dbName

	SysBioCol = SysBioClient.Database(SysBioNetDB).Collection(SysBioNetCol)

	err := SysBioCol.Database().RunCommand(context.TODO(), bson.D{{Key:"Ping", Value:1}}).Err()
	if err != nil {
		fmt.Printf("The Network DB %s will be created once we insert data in the collection %s \n",SysBioNetDB, SysBioNetCol)
	}
}

func VerifySysBiometics(c *gin.Context) {
	var sysBioMetrics model.SysBioMetrix

	/* when sending the struct from the local machine 
	   BindJSON - Reads the JSON payload from the request 
	   Mapping the JSON field to the Go struct the we have defind in the Azure vm models 
	*/
	if err := c.BindJSON(&sysBioMetrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// the global variable for the collection will be assigned by the below function
	dbColVerify(sysBioMetrics.SSID)
	
	existData, err := checkSysBioMetrics(sysBioMetrics)
	
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error checking data"})
		return
	}
	// the device biometrics is in the MongoDB, Biometrics
	if existData != nil {
		c.JSON(http.StatusOK, gin.H{"status":"The device biometrics in registered in the network db"})
	} else {
		if err := insertSysBioMetrics(sysBioMetrics); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failer to insert data"})
			return 
		}
		c.JSON(http.StatusOK, gin.H{"status" : "Data inserted"})
	}
}
	
func checkSysBioMetrics(sysBioMetrix model.SysBioMetrix) (*model.SysBioMetrix, error) {
	var result model.SysBioMetrix

	// Define filter to match the exact biometric data
	filter := bson.M{
		"ssid":                         sysBioMetrix.SSID,
		"mac":                          sysBioMetrix.MAC,
		"systemserialnumber":           sysBioMetrix.SystemSerialNumber,
		"uuid":                         sysBioMetrix.UUID,
	}
	
	// Query the testNetwork to check for existing record (Device Biometrix)
	err := SysBioCol.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No matching record found
		}
		return nil, err // Return any other error encountered
	}
	return &result, nil // Return the found record
}

func insertSysBioMetrics(sysBioMetrix model.SysBioMetrix) error {
	_, err := SysBioCol.InsertOne(context.TODO(), sysBioMetrix)
	return err
}

