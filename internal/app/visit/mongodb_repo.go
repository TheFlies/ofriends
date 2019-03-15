package visit

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
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Visit, error) {
	s := r.session.Clone()
	defer s.Close()
	var visit *types.Visit
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&visit); err != nil {
		return nil, errors.Wrap(err, "failed to find the given visit from database")
	}
	return visit, nil
}

// FindAll return all visits
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Visit, error) {
	s := r.session.Clone()
	defer s.Close()
	var visit []types.Visit
	if err := r.collection(s).Find(bson.M{}).All(&visit); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all visit from database")
	}
	
	return visit, nil
}

// Create a visit
func (r *MongoRepository) Create(ctx context.Context, visit types.Visit) error {
	s := r.session.Clone()
	defer s.Close()
	visit.ID = uuid.New()
	return r.collection(s).Insert(&visit)
}

// Update a visit
func (r *MongoRepository) Update(ctx context.Context, visit types.Visit) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"_id": visit.ID}, &visit)
}

// Delete a visit
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Remove(bson.M{"_id": id})
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("friends")
}
