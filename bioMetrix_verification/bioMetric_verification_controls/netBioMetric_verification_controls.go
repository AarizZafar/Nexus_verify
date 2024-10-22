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

const NetDBName = "Networks"
const NetColName = "NetbioMetrix"

var NetClient *mongo.Client
var NetCollection *mongo.Collection

func init() {
	NetClient = controllers.GetMongoSession()

	fmt.Println("\033[97;46m        ACCESS DB :", NetDBName, "COLLECTION :", NetColName, "           \033[0m")

	NetCollection = getNetCollection(NetClient)
	fmt.Println("\033[97;42m        SUCCESS IN ACCESSING COLLECTION : ", NetDBName, "       ✔      \033[0m")
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getNetCollection(client *mongo.Client) *mongo.Collection {
	NetCollection := client.Database(NetDBName).Collection(NetColName)

	/* RunCommand - executes a command on MongoDB {'ping' : 1} used to check the connection 
	   bson.D{{Key: "ping", Value: 1}} is a BSON (Binary JSON) document that MongoDB interprets as a “ping” request.
	*/
	err := NetCollection.Database().RunCommand(context.TODO(), bson.D{{Key:"ping", Value:1}}).Err()   // checks if the mongodb server is rechable by sending a simple "ping" command
	handleErr(err)

	if err != nil {
		// if the db does not exist the db will be CREATED automatically when we inster data into the DB
		fmt.Printf("the DataBase %s will be created when we insert data in the collection %s \n", NetDBName, NetColName)
	}
	return NetCollection
}

func VerifyNetBiometics(c *gin.Context) {
	var NetBioMetrics model.NetBioMetrix

	/* when sending the struct from the local machine 
	   BindJSON - Reads the JSON payload from the request 
	   Mapping the JSON field to the Go struct the we have defind in the Azure vm models 
	*/
	if err := c.BindJSON(&NetBioMetrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existData, err := checkNetBioMetrics(NetBioMetrics)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"ERROR CHECKING DATA"})
		return
	}

	if existData != nil {
		c.JSON(http.StatusOK, gin.H{"status":"THE NETWORK IS REGISTERES IN NETWORK DB"})
	} else {
		if err := insertNetBioMetrics(NetBioMetrics); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "FAILED TO INSERT NETWORK BIOMETRICS IN THE DB"})
			return 
		}
		c.JSON(http.StatusOK, gin.H{"status" : "NET BIOMETRIX INSERTED IN THE NETWORK DB"})
	}
}

func checkNetBioMetrics(NetBioMetrix model.NetBioMetrix) (*model.NetBioMetrix, error) {
	var result model.NetBioMetrix

	// Define filter to match the exact biometric data
	filter := bson.M{
		"ssid" 				: NetBioMetrix.SSID,
		"bssid" 			: NetBioMetrix.BSSID,
		"subnetmask" 		: NetBioMetrix.SubNetMask,
		"ipv4_dg" 			: NetBioMetrix.IPV4_DG,
		"ipv6_dg" 			: NetBioMetrix.IPV6_DG,
		"active_mac" 		: NetBioMetrix.Active_MAC,
		"inactive_mac" 		: NetBioMetrix.Inactive_MAC,
		"security_proto" 	: NetBioMetrix.Security_proto,
	}

	// Query the database to check for existing record
	err := NetCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No matching document found
		}
		return nil, err // Return any other error encountered
	}
	return &result, nil // Return the found document
}

func insertNetBioMetrics(NetBioMetrix model.NetBioMetrix) error {
	fmt.Println("\033[97;46m       INSERTING NETWORK BIOMETRIX IN NETWORK DB       \033[0m")
	_, err := NetCollection.InsertOne(context.TODO(), NetBioMetrix)
	return err
}
