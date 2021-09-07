package controllers

import (
	"context"
	"log"
	"net/http"

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
	c.FileAttachment("./assets/go/m.go", "m.go")
}
