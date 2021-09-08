package main

import (
	"github.com/prajwal-scorpionking123/DESTINATION_API/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("api/deployFiles", controllers.DeployeFiles)

	r.Static("../PRODUCTION", "../PRODUCTION")

	r.Run(":3002")
}
