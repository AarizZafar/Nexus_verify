package router

import (
	"github.com/AarizZafar/Nexus_verify.git/bioMetrix_verification"
	"github.com/AarizZafar/Nexus_verify.git/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)       // to prevent unecessry output to be shown on the terminal
	router := gin.Default() 

	router.LoadHTMLGlob("webPages/html/*")
	
	router.Static("webPages/css","./webPages/css")
	router.Static("webPages/js","./webPages/js")
	
	router.GET("/",     controllers.AdminLoginPage)

	// these 2 api will directly be connected to the biometric verification file
	router.POST("/NetVerify", bioMetrix_verification.InitiateNetVerification)
	router.POST("/SysVerify", bioMetrix_verification.InitiateSysVerification)

	router.GET("/GetSSIDS", controllers.GetSSIDS)
	
	router.GET("/GetAdmincreds", controllers.GetAdminCred)

	router.GET("/GetTestNetsfromSSID", controllers.GetTestNetsfromSSID)         // GETS ALL THE TEST NET FROM THE NETWORK, SSID

	router.GET("/GetBioMetrixtestNet", controllers.GetBioMetrixfromTestNet)     // GETS ALL THE BIOMETRIX FROM THE TEST NET
	
	router.POST("/login", controllers.AdminAuthentication)

	return router
}