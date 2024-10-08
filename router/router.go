package router

import (
	"github.com/gin-gonic/gin"
	"github.com/AarizZafar/Nexus_verify.git/controls"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)       // to prevent unecessry output to be shown on the terminal
	router := gin.Default() 

	router.POST("/verify", controls.VerifyBiometics)

	return router
}