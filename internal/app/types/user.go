package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	User struct {
		ID             string   `bson:"_id" json:"id"`
		Username       string   `bson:"username" json:"username"`
		FullName       string   `bson:"fullname" json:"fullname"`
		Email          string   `bson:"email" json:"email"`
		DeliveryCenter []string `bson:"delivery_center" json:"delivery_center"`
		Password       string   `bson:"password" json:"password"`
		OldPassword    string   `boson:"old_password"json:"old_password"`
		RoleName       string   `bson:"role" json:"role"`
		Priority       int      `bson:"priority"json:"priority"`
	}
)

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.Password, validation.NilOrNotEmpty),
		validation.Field(&u.FullName),
	)
}
