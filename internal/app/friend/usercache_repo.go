package friend

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

// NewUserCacheRepository is MongoDB implementation of repository
type UserCacheRepository struct {
	session *mgo.Session
}

// NewUserCacheRepository return new MongoDB repository
func NewUserCacheRepository(s *mgo.Session) *MongoRepository {
	return &MongoRepository{
		session: s,
	}
}

// FindByID return user in database base on given id
func (u *UserCacheRepository) FindUserByID(id string) (*types.Usercache, error) {
	s := u.session.Clone()
	defer s.Close()
	var user *types.Usercache
	if err := u.collection(s).Find(bson.M{"id": id}).One(&user); err != nil {
		return nil, errors.Wrap(err, id)
	}
	return user, nil
}
func (u *UserCacheRepository) InserUser(use *types.Usercache) error {
	s := u.session.Clone()
	defer s.Close()
	err := u.collection(s).Insert(&use)
	return err
}
func (u *UserCacheRepository) FindUserByUserName(username string) (*types.Usercache, error) {
	s := u.session.Clone()
	defer s.Close()
	var user *types.Usercache
	if err := u.collection(s).Find(bson.M{"username": username}).One(&user); err != nil {
		return nil, errors.Wrap(err, username)
	}
	return user, nil
}
func (u *UserCacheRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("usercache")
}
