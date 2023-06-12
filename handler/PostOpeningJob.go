package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabee/gopportunities/schemas"
)

func CreateOpeningJob(ctx *gin.Context) {
	request := RequestOpening{}

	ctx.BindJSON(&request)
	err := request.Validate()
	if err != nil {
		logger.Err("validation error: %v", err.Error())
	}
	opening := schemas.Opening{
		Salary:   request.Salary,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Remote:   *request.Remote,
		Role:     request.Role,
	}
	err = db.Create(&opening).Error
	if err != nil {
		logger.Err("error creating opening: %v", err.Error())
		return
	}
	SendSucess(ctx, "create-opening", opening)

}
