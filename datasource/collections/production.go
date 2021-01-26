package collections

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductionScheama struct {
	Name string `json:"name,omitempty"`
}

type Production struct {
	Client *mongo.Client
}

func (s *Production) FindAll() []ProductionScheama {
	var result []ProductionScheama
	collection := s.Client.Database(os.Getenv("MONGO_DB")).Collection("production")
	filter := bson.M{}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}
	return result
}
