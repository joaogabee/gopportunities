package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabee/gopportunities/schemas"
	"net/http"
)

func GetOpeningJobs(ctx *gin.Context) {
	opening := []schemas.Opening{}

	if err := db.Find(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	SendSucess(ctx, "list-opening", opening)
}
