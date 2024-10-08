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
	router.POST("/verify", bioMetrix_verification.InitiateVerification)

	return router
}