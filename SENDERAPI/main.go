package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prajwal-scorpionking123/SENDER/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/api/getSource", controllers.GetSources)
	router.GET("/api/sendFiles", controllers.SendFiles)
	router.Run(":3001")
}
