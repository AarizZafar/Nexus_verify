package router

import (
	"github.com/gin-gonic/gin"
	"github.com/AarizZafar/Nexus_verify.git/controls"
)

func Router() *gin.Engine {
	router := gin.Default() 

	router.POST("/verify", controls.VerifyBiometics)

	return router
}