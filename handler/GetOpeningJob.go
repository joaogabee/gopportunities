package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpeningJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
