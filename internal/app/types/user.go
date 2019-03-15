package types

import "github.com/globalsign/mgo/bson"

type (
	User struct {
		Id       bson.ObjectId `bson:"_id" json:"id"`
		Username string        `bson:"username" json:"username"`
		Fullname string        `bson:"fullname" json:"fullname"`
		Userrole string        `bson:"userrole" json:"userrole"`
		Email    string        `bson:"email" json:"email"`
	}
)
