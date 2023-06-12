package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":    msg,
		"Error Code": code,
	})
}
func SendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	})
}
