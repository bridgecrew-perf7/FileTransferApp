package controllers

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prajwal-scorpionking123/SENDER/helpers"
	"github.com/prajwal-scorpionking123/SENDER/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database = helpers.ConnectToMongoDB()
var collection *mongo.Collection = DB.Collection("Source")

func GetSources(c *gin.Context) {
	c.Header("content-type", "application/json")
	var sources []models.Source
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helpers.GetError(err, c)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var source models.Source
		err := cur.Decode(&source)
		if err != nil {
			log.Fatal(err)
		}

		sources = append(sources, source)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"sources": sources})
}
func SendFiles(c *gin.Context) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// c.FileAttachment("./assets/go/m.go", "m.go")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", "m.txt")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "failed",
		})
	}
	file, err := os.Open("./assets/go/m.txt")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "failed",
		})
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "failed",
		})
	}
	writer.Close()
	req, err := http.NewRequest("POST", "http://localhost:8080/sendfile", bytes.NewReader(body.Bytes()))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "failed",
		})
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "OK",
	})
}
