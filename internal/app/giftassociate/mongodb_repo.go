package giftassociate

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
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.GiftAssociate, error) {
	s := r.session.Clone()
	defer s.Close()
	var giftAssociate *types.GiftAssociate
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&giftAssociate); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return giftAssociate, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("ofriends").C("giftassociates")
}

// FindAll return all gift associates
func (r *MongoRepository) FindAll(ctx context.Context) ([]types.GiftAssociate, error) {
	s := r.session.Clone()
	defer s.Close()
	var giftAssociates []types.GiftAssociate
	if err := r.collection(s).Find(bson.M{}).All(&giftAssociates); err != nil {
		return nil, errors.Wrap(err, "failed to fetch all gifts from database")
	}

	return giftAssociates, nil
}

// FindByCusVisitAssocID return gift associate base on given ids
func (r *MongoRepository) FindByCusVisitAssocID(ctx context.Context, assignID string) ([]types.GiftAssociate, error) {
	s := r.session.Clone()
	defer s.Close()
	var giftAssociates []types.GiftAssociate
	if err := r.collection(s).Find(bson.M{"_custvisitassoc_id": assignID}).All(&giftAssociates); err != nil {
		return nil, errors.Wrap(err, "failed to find the given gift from database")
	}
	return giftAssociates, nil
}

// Create a gift associate
func (r *MongoRepository) Create(ctx context.Context, giftAssociate types.GiftAssociate) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	giftAssociate.ID = uuid.New()
	err := r.collection(s).Insert(&giftAssociate)
	return giftAssociate.ID, err
}

// Update a gift associate
func (r *MongoRepository) Update(ctx context.Context, giftAssociate types.GiftAssociate) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Update(bson.M{"_id": giftAssociate.ID}, &giftAssociate)
	return err
}

// Delete a gift associate
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_id": id})
	return err
}

// DeleteByCusVisitAssocID a gift associate
func (r *MongoRepository) DeleteByCusVisitAssocID(ctx context.Context, assignID string) error {
	s := r.session.Clone()
	defer s.Close()
	err := r.collection(s).Remove(bson.M{"_cusvisitassoc_id": assignID})
	return err
}

// UpdateNameByGiftID update gift name by gift id
func (r *MongoRepository) UpdateNameByGiftID(ctx context.Context, giftName string, giftID string) error {
	s := r.session.Clone()
	defer s.Close()
	filter := bson.M{"_gift_id": giftID}
	update := bson.M{"$set": bson.M{"giftname": giftName}}

	err := r.collection(s).Update(filter, update)
	return err
}

// IsAssignedGift check assigned gift
func (r *MongoRepository) IsAssignedGift(ctx context.Context, giftID string, cusvisitassocID string) bool {
	s := r.session.Clone()
	defer s.Close()
	count, err := r.collection(s).Find(bson.M{
		"$or": []bson.M{
			bson.M{"_gift_id": giftID},
			bson.M{"_cusvisitassoc_id": cusvisitassocID},
		},
	}).Limit(1).Count()
	if err == nil && count > 0 {
		return true
	}
	return false
}
