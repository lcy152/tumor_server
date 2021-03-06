package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApplication struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Friend      primitive.ObjectID `json:"friend" bson:"friend"`
	FriendName  string             `json:"friend_name" bson:"friend_name"`
	Status      string             `json:"status" bson:"status"`
	User        primitive.ObjectID `json:"user" bson:"user"`
	UserName    string             `json:"user_name" bson:"user_name"`
	Description string             `json:"description" bson:"description"`
	CreateTime  time.Time          `json:"create_time" bson:"create_time"`
	OperateTime time.Time          `json:"operate_time" bson:"operate_time"`
}
