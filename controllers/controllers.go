package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLoginPage(ctx *gin.Context) {
	fmt.Println("----------------Admin Login Page-------------------")
	ctx.HTML(http.StatusOK, "login_page_admin.html", nil)
}