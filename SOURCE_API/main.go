package main

import (
	"github.com/gin-gonic/gin"
	"github.com/team_six/SOURCE_API/controllers"
	"github.com/team_six/SOURCE_API/controllers/authcontroller"
)

func main() {
	router := gin.Default()
	router.Static("../SOURCE", "../SOURCE")
	router.POST("/api/postlink", controllers.PostLink)
	router.GET("/api/getSources", controllers.GetSources)
	router.POST("/api/deployFiles", controllers.DeployFiles)
	router.POST("/login", authcontroller.Auth)
	router.Run(":3001")
}
