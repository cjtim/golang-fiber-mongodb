package collections

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserScheama struct {
	name    string
	lineUid string
}

type User struct {
	Client *mongo.Client
}

func (s *User) FindAll() []UserScheama {
	var result []UserScheama
	collection := s.Client.Database(os.Getenv("MONGO_DB")).Collection("users")
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
