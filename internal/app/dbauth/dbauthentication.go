package dbauth

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/TheFlies/ofriends/internal/app/api/handler/user"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
)

type (
	DBAuthentication struct {
		log     glog.Logger
		userSrv user.UserService
	}
)

func NewDBAuthentication(l glog.Logger, userSrv user.UserService) *DBAuthentication {
	return &DBAuthentication{
		log:     l,
		userSrv: userSrv,
	}
}
func (dba *DBAuthentication) Authenticate(username string, password string) (interface{}, error) {
	dba.log.Infof("%v is authenticated by database ", username)
	user, err := dba.userSrv.GetByName(username)
	if err != nil {
		return nil, errors.Wrap(err, "user is not found")
	}
	dbPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		dba.log.Errorf("can't compare two password, the input password isn't match with database  err: %v", err)
		return nil, errors.Wrap(err, "wrong password ")
	}
	dba.log.Infof("db login successfully")
	returnUser := types.User{
		ID:             user.ID,
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		DeliveryCenter: user.DeliveryCenter,
	}
	tokenGeneration := jwt.NewTokenGeneration()
	jwtToken, err := tokenGeneration.CreateToken(user.Username, user.FullName, user.DeliveryCenter)
	if err != nil {
		dba.log.Errorf("can't create token form user information err: %v ", err)
		return nil, errors.Wrap(err, "create token fail ")
	}
	//remove hashed password
	returnUser.Password = ""
	respondMap := make(map[string]interface{})
	respondMap["userInformation"] = returnUser
	respondMap["token"] = jwtToken
	return respondMap, nil
}
