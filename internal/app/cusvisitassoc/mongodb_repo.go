package cusvisitassoc

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
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.CusVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var cusVisitAssoc *types.CusVisitAssoc
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&cusVisitAssoc); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return cusVisitAssoc, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("cusvisitassocs")
}

// FindAll return all activity visit associates
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.CusVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var cusVisitAssocs []types.CusVisitAssoc
	if err := r.collection(s).Find(bson.M{}).All(&cusVisitAssocs); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all gifts from database")
	}

	return cusVisitAssocs, nil
}

// FindByVisitID return activity visit associate base on given ids
func (r *MongoRepository) FindByVisitID(ctx context.Context, visitID string) ([]types.CusVisitAssoc, error) {
	s := r.session.Clone()
	defer s.Close()
	var cusVisitAssocs []types.CusVisitAssoc
	if err := r.collection(s).Find(bson.M{"_visit_id": visitID}).All(&cusVisitAssocs); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return cusVisitAssocs, nil
}

// Create a activity visit associate
func (r *MongoRepository) Create(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	cusVisitAssoc.ID = uuid.New()
	err := r.collection(s).Insert(&cusVisitAssoc)
	return cusVisitAssoc.ID, err
}

// Update a activity visit associate
func (r *MongoRepository) Update(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"_id": cusVisitAssoc.ID}, &cusVisitAssoc)
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

// UpdateNameByCusID update customer name by customer id
func (r *MongoRepository) UpdateNameByCusID(ctx context.Context, customerName string, customerID string) error {
	s := r.session.Clone()
	defer s.Close()
	filter := bson.M{"_customer_id": customerID}
	update := bson.M{"$set": bson.M{"customername": customerName}}

	err := r.collection(s).Update(filter, update)
	return err
}

// IsAssignedCustomer check assigned customer
func (r *MongoRepository) IsAssignedCustomer(ctx context.Context, customerID string, visitID string) bool {
	s := r.session.Clone()
	defer s.Close()
	count, err := r.collection(s).Find(bson.M{
		"$or": []bson.M{
			bson.M{"_customer_id": customerID},
			bson.M{"_visit_id": visitID},
		},
	}).Limit(1).Count()
	fmt.Println(count)
	if err == nil && count > 0 {
		return true
	}
	return false
}
