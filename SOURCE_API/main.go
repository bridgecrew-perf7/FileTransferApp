package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prajwal-scorpionking123/SOURCE_API/controllers"
)

func main() {
	router := gin.Default()
	router.Static("/SOURCE", "./SOURCE")
	router.POST("/api/postlink", controllers.PostLink)
	router.GET("/api/getSources", controllers.GetSources)
	router.POST("/api/deployFiles", controllers.DeployFiles)
	router.Run(":3001")
}
