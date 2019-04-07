package gift

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/uuid"
	"github.com/globalsign/mgo"
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

// FindByID return gift base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Gift, error) {
	s := r.session.Clone()
	defer s.Close()
	var gift *types.Gift
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&gift); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return gift, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("gifts")
}

// FindAll return all gifts
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Gift, error) {
	s := r.session.Clone()
	defer s.Close()
	var gifts []types.Gift
	if err := r.collection(s).Find(bson.M{}).All(&gifts); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all gifts from database")
	}

	return gifts, nil
}

// Create a gift
func (r *MongoRepository) Create(ctx context.Context, gift types.Gift) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	gift.ID = uuid.New()
	err := r.collection(s).Insert(&gift)
	return gift.ID, err
}

// Update a gift
func (r *MongoRepository) Update(ctx context.Context, gift types.Gift) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"_id": gift.ID}, &gift)
	return err
}

// Delete a gift
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_id": id})
	return err
}
