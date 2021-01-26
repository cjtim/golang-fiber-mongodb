package models

import (
	"github.com/cjtim/golang-fiber-mongodb/datasource/collections"
	"go.mongodb.org/mongo-driver/mongo"
)

// Models -- center for invoke any operation on Model
type Models struct {
	Production collections.Production
	User       collections.User
}

// GetModels by pass Mongo Client pointer
func GetModels(c *mongo.Client) *Models {
	models := &Models{}
	models.Production.Client = c
	models.User.Client = c
	return models
}
