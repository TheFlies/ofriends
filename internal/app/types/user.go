package types

import "github.com/globalsign/mgo/bson"

type (
	User struct {
		Id             bson.ObjectId `bson:"_id" json:"id"`
		Username       string        `bson:"username" json:"username"`
		Fullname       string        `bson:"fullname" json:"fullname"`
		Email          string        `bson:"email" json:"email"`
		DeliveryCenter string        `bson:"delivery_center" json:"delivery_center"`
	}
)
