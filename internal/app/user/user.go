package user

import (
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/TheFlies/ofriends/internal/app/ldap"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/jwt"
)

type (
	UserReponsitory interface {
		FindUserByUserName(username string) (*types.Usercache, error)
		InserUser(user *types.Usercache) error
		CheckUserByUsername(username string) bool
	}
	UserService struct {
		repo   UserReponsitory
		logger glog.Logger
	}
)

func NewUserCacheService(r UserReponsitory, l glog.Logger) *UserService {
	return &UserService{
		repo:   r,
		logger: l,
	}
}
func (s *UserService) GetByName(username string) (*types.Usercache, error) {
	return s.repo.FindUserByUserName(username)
}
func (s *UserService) AddUser(user *types.Usercache) error {
	return s.repo.InserUser(user)
}
func (s *UserService) CheckExistence(username string) bool {
	return s.repo.CheckUserByUsername(username)
}
func (s *UserService) LDAPAuthenticate(username string, password string) (interface{}, error) {
	err := ldap.LDAPAuthentication(username, password)
	if err != nil {
		s.logger.Errorf("can't authenticate with username: %s ", username)
		newerror := errors.Wrap(err, "Can't authenticate by username")
		return nil, newerror
	}
	respondmap := make(map[string]interface{})
	dbuser, err := s.repo.FindUserByUserName(username)
	if err != nil {
		s.logger.Infof("can't find username: %s in database, the information will be finded in ldap database", username)
		ldapuser, err := ldap.LDAPQuery(username)
		if err != nil {
			s.logger.Errorf("error when query user information in ldap server")
			return nil, errors.Wrap(err, "error when query ldap information")
		}
		ldapuser.Id = bson.NewObjectId()
		err = s.repo.InserUser(&ldapuser)
		if err != nil {
			s.logger.Errorf("%v,can't insert user into database", err)
			return nil, errors.Wrap(err, "can't insert user into database")
		}
		jwttoken, err := jwt.CreateToken(ldapuser.Username, ldapuser.Fullname, ldapuser.Userrole)
		if err != nil {
			s.logger.Errorf("%v,creating token is fails", err)
			return nil, errors.Wrap(err, "creating token is fails")
		}

		respondmap["userinformation"] = ldapuser
		respondmap["token"] = jwttoken
		return respondmap, nil
	}
	jwttoken, err := jwt.CreateToken(dbuser.Username, dbuser.Fullname, dbuser.Userrole)
	if err != nil {
		s.logger.Errorf("%v,creating token is fails", err)
		return nil, errors.Wrap(err, "creating token is fails")
	}
	respondmap["userinformation"] = dbuser
	respondmap["token"] = jwttoken
	return respondmap, nil
}

//TODO this fuction will be implemented later
//func (service *UserService)DatabaseAuthenticate(username string,password string)bool {
//
//}
