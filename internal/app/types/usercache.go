package types

type (
	Usercache struct {
		Username string `json:"username" bson:"username"`
		Fullname string `json:"name"`
		Userrole string `json:"userrole"`
	}
)
