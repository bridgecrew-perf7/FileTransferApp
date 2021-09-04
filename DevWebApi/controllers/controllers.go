package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prajwal-scorpionking123/DevWebApi/helpers"
	"github.com/prajwal-scorpionking123/DevWebApi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database = helpers.ConnectToMongoDB()
var collection *mongo.Collection = DB.Collection("Source")

func PostLink(c *gin.Context) {

	c.Header("content-type", "application/json")

	var source models.Source

	if err := c.Bind(&source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	source.Timestamp = time.Now()
	res, err := collection.InsertOne(context.TODO(), source)
	if err != nil {
		helpers.GetError(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"source": res})
}
func GetLinks(c *gin.Context) {
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
func GetLink(c *gin.Context) {
	c.Header("content-type", "application/json")
	var source models.Source

	var params = c.Param("id")

	id, _ := primitive.ObjectIDFromHex(params)

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&source)

	if err != nil {
		helpers.GetError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"source": source})
}
