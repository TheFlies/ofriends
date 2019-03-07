package usercache

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

type (
	UsercacheMongoReponsitory struct {
		session *mgo.Session
	}
)

func NewUsercachMongoReopnsitoty(s *mgo.Session) *UsercacheMongoReponsitory {
	return &UsercacheMongoReponsitory{
		session: s,
	}
}
func (r *UsercacheMongoReponsitory) FindUserByUserName(username string) (*types.Usercache, error) {
	s := r.session.Clone()
	defer s.Close()
	var user *types.Usercache
	if err := r.collection(s).Find(bson.M{"username": username}).One(&user); err != nil {
		return nil, errors.Wrap(err, "Fail to find username in database")
	}
	return user, nil
}
func (r *UsercacheMongoReponsitory) InserUser(user *types.Usercache) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Insert(&user); err != nil {
		return errors.Wrap(err, "Fail insert user to database")
	}
	return nil

}
func (r *UsercacheMongoReponsitory) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("usercache")
}
