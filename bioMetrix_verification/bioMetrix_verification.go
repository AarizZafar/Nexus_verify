package bioMetrix_verification

import (
	"github.com/gin-gonic/gin"
	"github.com/AarizZafar/Nexus_verify.git/bioMetrix_verification/bioMetric_verification_controls"
)

func InitiateVerification(ctx *gin.Context) {
	bioMetric_verification_controls.VerifyBiometics(ctx)
}