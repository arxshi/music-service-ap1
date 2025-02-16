// /models/album.go

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AlbumCollection *mongo.Collection

type Album struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	Title  string               `bson:"title"`
	Artist string               `bson:"artist"`
	Tracks []primitive.ObjectID `bson:"tracks"`
}
