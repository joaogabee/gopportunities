package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabee/gopportunities/schemas"
	"net/http"
)

func PutOpeningJob(ctx *gin.Context) {
	request := UpdateOpening{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Err("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")

	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "opening not found")
		return
	}
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Err("error updating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	SendSucess(ctx, "update opening", opening)
}
