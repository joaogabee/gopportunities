package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabee/gopportunities/schemas"
	"net/http"
)

func GetOpeningJob(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameters is required").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "opening not found")
	}
	SendSucess(ctx, "show-opening", opening)
}
