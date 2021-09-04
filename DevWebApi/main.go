package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prajwal-scorpionking123/DevWebApi/controllers"
)

func main() {
	router := gin.New()
	router.POST("/api/postlink", controllers.PostLink)
	router.GET("/api/getlinks", controllers.GetLinks)
	router.GET("/api/getLink", controllers.GetLinks)

	router.Run(":3000")
}
