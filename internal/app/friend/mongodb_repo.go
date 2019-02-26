package friend

import (
	"context"

	"ofriends/internal/app/types"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

// MongoRepository is MongoDB implementation of repository
type MongoRepository struct {
	session *mgo.Session
}

// NewMongoRepository return new MongoDB repository
func NewMongoRepository(s *mgo.Session) *MongoRepository {
	return &MongoRepository{
		session: s,
	}
}

// FindByID return friend base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Friend, error) {
	s := r.session.Clone()
	defer s.Close()
	var friend *types.Friend
	if err := r.collection(s).Find(bson.M{"id": id}).One(&friend); err != nil {
		return nil, errors.Wrap(err, "failed to find the given friend from database")
	}
	return friend, nil
}

// FindAll return all friends
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Friend, error) {
	s := r.session.Clone()
	defer s.Close()
	var friends []types.Friend
	if err := r.collection(s).Find(bson.M{}).All(&friends); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all friends from database")
	}
	return friends, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("friends").C("friends")
}
