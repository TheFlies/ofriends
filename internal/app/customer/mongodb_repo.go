package customer

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

// FindByID return customer base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Customer, error) {
	s := r.session.Clone()
	defer s.Close()
	var customer *types.Customer
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&customer); err != nil {
		return nil, errors.Wrap(err, "failed to find the given customer from database")
	}
	return customer, nil
}

// FindAll return all customers
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.Customer, error) {
	s := r.session.Clone()
	defer s.Close()
	var customer []types.Customer
	if err := r.collection(s).Find(bson.M{}).All(&customer); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all customer from database")
	}
	
	return customer, nil
}

// Create a customer
func (r *MongoRepository) Create(ctx context.Context, customer types.Customer) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	customer.ID = uuid.New()
	err := r.collection(s).Insert(&customer)
	return customer.ID, err
}

// Update a customer
func (r *MongoRepository) Update(ctx context.Context, customer types.Customer) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"_id": customer.ID}, &customer)
}

// Delete a customer
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Remove(bson.M{"_id": id})
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("customers")
}
