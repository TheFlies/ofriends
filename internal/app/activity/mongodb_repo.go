package activity

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

// FindByID return activity base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Activity, error) {
	s := r.session.Clone()
	defer s.Close()
	var act *types.Activity
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&act); err != nil {
		return nil, errors.Wrap(err, "failed to find the given activity from database")
	}
	return act, nil
}

// Find all activities by visit ID
func (r *MongoRepository) FindByVisitID(ctx context.Context, visitId string) ([]types.Activity, error) {
	s := r.session.Clone()
	defer s.Close()
	var act []types.Activity
	if err := r.collection(s).Find(bson.M{"_visit_id": visitId}).All(&act); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all activities from database")
	}
	
	return act, nil
}

// FindAll return all activites
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Activity, error) {
	s := r.session.Clone()
	defer s.Close()
	var act []types.Activity
	if err := r.collection(s).Find(bson.M{}).All(&act); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all activity from database")
	}
	
	return act, nil
}

// Create a activity
func (r *MongoRepository) Create(ctx context.Context, act types.Activity) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	act.ID = uuid.New()
	err := r.collection(s).Insert(&act)
	return act.ID, err
}

// Update a activity
func (r *MongoRepository) Update(ctx context.Context, act types.Activity) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"_id": act.ID}, &act)
}

// Delete a activity
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Remove(bson.M{"_id": id})
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("activities")
}
