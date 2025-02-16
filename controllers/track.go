// /controllers/track.go

package controllers

import (
	"AP1/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetTracks(c *gin.Context, collection *mongo.Collection) {
	var tracks []models.Track
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tracks"})
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tracks"})
		}
	}(cursor, context.Background())
	for cursor.Next(context.Background()) {
		var track models.Track
		if err := cursor.Decode(&track); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding track"})
			return
		}
		tracks = append(tracks, track)
	}
	c.JSON(http.StatusOK, tracks)
}

func GetTrack(c *gin.Context, collection *mongo.Collection) {
	trackID, err := primitive.ObjectIDFromHex(c.Param("track_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track ID"})
		return
	}
	var track models.Track
	err = collection.FindOne(context.Background(), bson.M{"_id": trackID}).Decode(&track)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}
	c.JSON(http.StatusOK, track)
}

func UploadTrack(c *gin.Context, collection *mongo.Collection) {
	var track models.Track
	if err := c.BindJSON(&track); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	track.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), track)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload track"})
		return
	}
	c.JSON(http.StatusOK, track)
}

func DeleteTrack(c *gin.Context, collection *mongo.Collection) {
	trackID, _ := primitive.ObjectIDFromHex(c.Param("track_id"))
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": trackID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete track"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Track deleted successfully"})
}
