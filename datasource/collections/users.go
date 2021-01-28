package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserScheama struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	LineUid string             `json:"lineUid,omitempty" bson:"lineUid,omitempty"`
}
