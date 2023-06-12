package router

import (
	"github.com/gin-gonic/gin"
)

func Initializer() {
	router := gin.Default()
	router.Run(":8080")
}
