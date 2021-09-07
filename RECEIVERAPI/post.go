package main

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	MyFile *multipart.FileHeader `form:"myfile" binding:"required"`
}

func main() {
	r := gin.Default()
	// r.GET("")
	r.POST("/sendfile", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBind(&userObj); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		err := c.SaveUploadedFile(userObj.MyFile, "assets/"+userObj.MyFile.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	r.Static("assets", "./assets")

	r.Run("localhost:8080")
}
