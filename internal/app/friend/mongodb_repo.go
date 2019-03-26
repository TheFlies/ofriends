package friend

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/uuid"
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
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&friend); err != nil {
		return nil, errors.Wrap(err, "failed to find the given friend from database")
	}
	return friend, nil
}

// FindAll return all friends
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Friend, error) {
	s := r.session.Clone()
	defer s.Close()
	var friend []types.Friend
	if err := r.collection(s).Find(bson.M{}).All(&friend); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all friend from database")
	}
	
	return friend, nil
}

// Create a friend
func (r *MongoRepository) Create(ctx context.Context, friend types.Friend) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	friend.ID = uuid.New()
	err := r.collection(s).Insert(&friend)
	return friend.ID, err
}

// Update a friend
func (r *MongoRepository) Update(ctx context.Context, friend types.Friend) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"_id": friend.ID}, &friend)
}

// Delete a friend
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Remove(bson.M{"_id": id})
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("friends")
}
