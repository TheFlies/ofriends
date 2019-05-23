package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/uuid"
)

type (
	UserMongoRepository struct {
		session *mgo.Session
	}
)

func NewUserMongoRepository(s *mgo.Session) *UserMongoRepository {
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
func (r *UserMongoRepository) InsertUser(u *types.User) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	u.ID = uuid.New()
	if err := r.collection(s).Insert(&u); err != nil {
		return "", errors.Wrap(err, "insert u to database is fail")
	}
	return u.Username, nil
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
	return s.DB("").C("users")
}
func (r *UserMongoRepository) UpdateUser(u *types.User) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"username": u.Username}, u)
	if err != nil {
		return errors.Wrap(err, "can't update u")
	}
	return nil
}
