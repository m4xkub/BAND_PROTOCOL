package controller

import (
	"ThridQuestion/internal/service"

	"github.com/gin-gonic/gin"
)

func GetRootController() *gin.Engine {
	router := gin.Default()

	router.POST("/CreateTransaction", service.BoardCast)

	return router
}
