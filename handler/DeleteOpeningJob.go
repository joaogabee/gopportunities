package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joaogabee/gopportunities/schemas"
	"net/http"
)

func DeleteOpeningJob(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening ID: %s", id))
		return
	}
	SendSucess(ctx, "delete-opening", opening)
}
