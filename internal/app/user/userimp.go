package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/TheFlies/ofriends/internal/app/types"
)

type (
	UserMongoRepository struct {
		session *mgo.Session
	}
)

func NewUserMongoRepositoty(s *mgo.Session) *UserMongoRepository {
	return &UserMongoRepository{
		session: s,
	}
}
func (r *UserMongoRepository) FindUserByUserName(username string) (*types.User, error) {
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(bson.M{"username": username}).One(&user); err != nil {
		return nil, errors.Wrap(err, "find username in database is fails")
	}
	return user, nil
}
func (r *UserMongoRepository) InserUser(user *types.User) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Insert(&user); err != nil {
		return errors.Wrap(err, "insert user to database is fail")
	}
	return nil

}
func (r *UserMongoRepository) CheckUserByUsername(username string) bool {
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(bson.M{"username": username}).One(&user); err != nil {
		return false
	}
	return true
}
func (r *UserMongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("user")
}
func (r *UserMongoRepository) UpdateUser(user *types.User) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"username": user.Username}, user)
	if err != nil {
		return errors.Wrap(err, "can't update user")
	}
	return nil
}
