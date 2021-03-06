package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddInstitutionApplication struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	User        primitive.ObjectID `json:"user" bson:"user"`
	UserName    string             `json:"user_name" bson:"user_name"`
	Status      string             `json:"status" bson:"status"`
	Institution *Institution       `json:"institution" bson:"institution"`
	Description string             `json:"description" bson:"description"`
	CreateTime  time.Time          `json:"create_time" bson:"create_time"`
	OperateTime time.Time          `json:"operate_time" bson:"operate_time"`
}
