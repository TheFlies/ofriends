package actvisitassoc

import (
	"context"
	"fmt"

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

// FindByID return activity visit associate base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.ActVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var actVisitAssoc *types.ActVisitAssoc
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&actVisitAssoc); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return actVisitAssoc, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("actvisitassocs")
}

// FindAll return all activity visit associates
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.ActVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var actVisitAssocs []types.ActVisitAssoc
	if err := r.collection(s).Find(bson.M{}).All(&actVisitAssocs); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all gifts from database")
	}

	return actVisitAssocs, nil
}

// FindByVisitID return activity visit associate base on given ids
func (r *MongoRepository) FindByVisitID(ctx context.Context, visitID string) ([]types.ActVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var actVisitAssocs []types.ActVisitAssoc
	if err := r.collection(s).Find(bson.M{"_visit_id": visitID}).All(&actVisitAssocs); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return actVisitAssocs, nil
}

// FindByActID return activity visit associate base on given ids
func (r *MongoRepository) FindByActID(ctx context.Context, actID string) ([]types.ActVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var actVisitAssocs []types.ActVisitAssoc
	if err := r.collection(s).Find(bson.M{"_activity_id": actID}).All(&actVisitAssocs); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return actVisitAssocs, nil
}

// Create a activity visit associate
func (r *MongoRepository) Create(ctx context.Context, actVisitAssoc types.ActVisitAssoc) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	actVisitAssoc.ID = uuid.New()
	err := r.collection(s).Insert(&actVisitAssoc)
	return actVisitAssoc.ID, err
}

// Update a activity visit associate
func (r *MongoRepository) Update(ctx context.Context, actVisitAssoc types.ActVisitAssoc) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"_id": actVisitAssoc.ID}, &actVisitAssoc)
	return err
}

// Delete a activity visit associate
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_id": id})
	return err
}

// DeleteByVisitID a activity visit associate
func (r *MongoRepository) DeleteByVisitID(ctx context.Context, visitID string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_visit_id": visitID})
	return err
}

// UpdateNameByActID update activity name by activity id
func (r *MongoRepository) UpdateNameByActID(ctx context.Context, activityName string, activityID string) error {
	s := r.session.Clone()
	defer s.Close()
	filter := bson.M{"_activity_id": activityID}
	update := bson.M{"$set": bson.M{"activityname": activityName}}

	err := r.collection(s).Update(filter, update)
	return err
}

// IsAssignedActivity check assigned activity
func (r *MongoRepository) IsAssignedActivity(ctx context.Context, activityID string, visitID string) bool {
	s := r.session.Clone()
	defer s.Close()
	count, err := r.collection(s).Find(bson.M{
		"$or": []bson.M{
			bson.M{"_activity_id": activityID},
			bson.M{"_visit_id": visitID},
		},
	}).Limit(1).Count()
	fmt.Println(count)
	if err == nil && count > 0 {
		return true
	}
	return false
}
