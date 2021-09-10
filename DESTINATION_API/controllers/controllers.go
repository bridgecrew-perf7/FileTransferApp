package controllers

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/team_six/DESTINATION_API/models"
)

//deploy single file controller
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

//fucion for taking the backup
func BackupFiles(c *gin.Context) {
	ZipWriter()
}

//helper function here
func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
func ZipWriter() {
	baseFolder := "../PRODUCTION/"
	output := "../BACKUP/bamu"

	err := os.MkdirAll(output, 0755)
	// Get a Buffer to Write To
	if err != nil {
		fmt.Println(err)
	}
	outFile, err := os.Create("../BACKUP/bamu/dubackup.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, baseFolder, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}
