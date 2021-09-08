package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/prajwal-scorpionking123/DESTINATION_API/models"
)

func DeployeFiles(c *gin.Context) {
	var fileMeta models.FileMeta
	if err := c.ShouldBind(&fileMeta); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	fileName := filepath.Base(fileMeta.File.Filename)
	println(fileName)
	err := c.SaveUploadedFile(fileMeta.File, "../PRODUCTION/"+fileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   fileMeta,
	})
}
