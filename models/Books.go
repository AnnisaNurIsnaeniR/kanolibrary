package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Books struct {
	Id  	primitive.ObjectID `bson:"_id"`
	Title 	string `bson:"title"`
	Author	string `bson:"author"`
	Publisher string `bson:"publisher"`
	Synopsis string `bson:"synopsis"`
}