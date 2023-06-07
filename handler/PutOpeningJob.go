package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutOpeningJob(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
