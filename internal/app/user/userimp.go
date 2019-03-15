package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/TheFlies/ofriends/internal/app/types"
)

type (
	UserMongoReponsitory struct {
		session *mgo.Session
	}
)

func NewUserMongoReopnsitoty(s *mgo.Session) *UserMongoReponsitory {
	return &UserMongoReponsitory{
		session: s,
	}
}
func (r *UserMongoReponsitory) FindUserByUserName(username string) (*types.User, error) {
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(bson.M{"username": username}).One(&user); err != nil {
		return nil, errors.Wrap(err, "fail to find username in database")
	}
	return user, nil
}
func (r *UserMongoReponsitory) InserUser(user *types.User) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Insert(&user); err != nil {
		return errors.Wrap(err, "fail insert user to database")
	}
	return nil

}
func (r *UserMongoReponsitory) CheckUserByUsername(username string) bool {
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(bson.M{"username": username}).One(&user); err != nil {
		return false
	}
	return true
}
func (r *UserMongoReponsitory) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("user")
}
