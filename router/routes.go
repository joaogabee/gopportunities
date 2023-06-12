package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabee/gopportunities/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/openings", handler.GetOpeningJobs)

		v1.GET("/opening", handler.GetOpeningJob)
		v1.POST("/opening", handler.CreateOpeningJob)
		v1.DELETE("/opening/{id}", handler.DeleteOpeningJob)
		v1.PUT("/opening", handler.PutOpeningJob)
	}
}
