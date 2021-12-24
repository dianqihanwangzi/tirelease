package api

import (
	"tirelease/internal/service"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Create gin-routers
func Routers(file string) (router *gin.Engine) {
	router = gin.New()

	// Static html
	router.Use(
		static.Serve("/", static.LocalFile(file, true)),
	)

	// Test "ping"
	ping := router.Group("/ping")
	{
		ping.GET("/", pong)
	}

	// REST API registry
	testEntity := router.Group("/testentity")
	{
		testEntity.POST("/insert", service.TestEntityInsert)
	}

	return router
}
