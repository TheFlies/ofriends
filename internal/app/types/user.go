package types

import "github.com/globalsign/mgo/bson"

type (
	User struct {
		ID             bson.ObjectId `bson:"_id" json:"id"`
		Username       string        `bson:"username" json:"username"`
		FullName       string        `bson:"fullname" json:"fullname"`
		Email          string        `bson:"email" json:"email"`
		DeliveryCenter []string      `bson:"delivery_center" json:"delivery_center"`
		Password       string        `bson:"password" json:"password"`
	}
)
