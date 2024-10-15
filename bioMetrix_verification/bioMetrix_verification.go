package bioMetrix_verification

import (
	"github.com/gin-gonic/gin"
	"github.com/AarizZafar/Nexus_verify.git/bioMetrix_verification/bioMetric_verification_controls"
)

func InitiateNetVerification(ctx *gin.Context) {
	bioMetric_verification_controls.VerifyNetBiometics(ctx)
}

func InitiateSysVerification(ctx *gin.Context) {
	bioMetric_verification_controls.VerifySysBiometics(ctx)
}