package activityvisitassociate

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

// FindByID return gift associate base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.ActivityVisitAssociate, error) {
	s := r.session.Clone()
	defer s.Close()
	var activityVisitAssociate *types.ActivityVisitAssociate
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&activityVisitAssociate); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return activityVisitAssociate, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("activityVisitAssociate")
}

// FindAll return all gift associates
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.ActivityVisitAssociate, error) {
	s := r.session.Clone()
	defer s.Close()
	var activityVisitAssociates []types.ActivityVisitAssociate
	if err := r.collection(s).Find(bson.M{}).All(&activityVisitAssociates); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all gifts from database")
	}

	return activityVisitAssociates, nil
}

// FindByVisitID return gift associate base on given ids
func (r *MongoRepository) FindByVisitID(ctx context.Context, visitID string) ([]types.ActivityVisitAssociate, error) {
	s := r.session.Clone()
	defer s.Close()
	var activityVisitAssociates []types.ActivityVisitAssociate
	if err := r.collection(s).Find(bson.M{"_visit_id": visitID}).All(&activityVisitAssociates); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return activityVisitAssociates, nil
}

// Create a gift associate
func (r *MongoRepository) Create(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	activityVisitAssociate.ID = uuid.New()
	err := r.collection(s).Insert(&activityVisitAssociate)
	return activityVisitAssociate.ID, err
}

// Update a gift associate
func (r *MongoRepository) Update(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"_id": activityVisitAssociate.ID}, &activityVisitAssociate)
	return err
}

// Delete a gift associate
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_id": id})
	return err
}

// DeleteByVisitID a gift associate
func (r *MongoRepository) DeleteByVisitID(ctx context.Context, visitID string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_visit_id": visitID})
	return err
}
