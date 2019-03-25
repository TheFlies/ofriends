package dbauth

import (
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/TheFlies/ofriends/internal/app/api/handler/login"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
)

type (
	EncryptionConf struct {
		Cost int
	}
	DBAuthentication struct {
		log     glog.Logger
		userSrv login.UserService
	}
)

func NewDBAuthentication(l glog.Logger, usersrv login.UserService) *DBAuthentication {
	return &DBAuthentication{
		log:     l,
		userSrv: usersrv,
	}
}
func (dba *DBAuthentication) Authenticate(username string, password string) (interface{}, error) {
	dba.log.Infof("%v is authenticated by database ", username)
	user, err := dba.userSrv.GetByName(username)
	if err != nil {
		return nil, errors.Wrap(err, "user is not found")
	}
	DBPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(password))
	if err != nil {
		dba.log.Errorf("can't compare two password, the input password isn't match with database  err: %v", err)
		return nil, errors.Wrap(err, "wrong password ")
	}
	dba.log.Infof("db login successfully")
	returnUser := types.User{
		ID:             bson.NewObjectId(),
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		DeliveryCenter: user.DeliveryCenter,
	}
	tokenGeneration := jwt.NewTokenGeneration()
	JWTToken, err := tokenGeneration.CreateToken(user.Username, user.FullName, user.DeliveryCenter)
	if err != nil {
		dba.log.Errorf("can't create token form user information err: %v ", err)
		return nil, errors.Wrap(err, "create token fail ")
	}
	respondMap := make(map[string]interface{})
	respondMap["userInformation"] = returnUser
	respondMap["token"] = JWTToken
	return respondMap, nil
}
