package giftassociate

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a gift repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.GiftAssociate, error)
	FindAll(ctx context.Context) ([]types.GiftAssociate, error)
	FindByVisitID(ctx context.Context, visitID string) ([]types.GiftAssociate, error)
	Create(ctx context.Context, giftAssociate types.GiftAssociate) (string, error)
	Update(ctx context.Context, giftAssociate types.GiftAssociate) error
	Delete(ctx context.Context, id string) error
}

// Service is an gift associate service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new gift associate service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given gift associate by id
func (s *Service) Get(ctx context.Context, id string) (*types.GiftAssociate, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAll return all gift associates from database
func (s *Service) GetAll(ctx context.Context) ([]types.GiftAssociate, error) {
	return s.repo.FindAll(ctx)
}

// GetByVisitID return given gift associate by visit id
func (s *Service) GetByVisitID(ctx context.Context, visitID string) ([]types.GiftAssociate, error) {
	return s.repo.FindByVisitID(ctx, visitID)
}

// Create a gift associates
func (s *Service) Create(ctx context.Context, giftAssociate types.GiftAssociate) (string, error) {
	if err := validation.ValidateStruct(&giftAssociate,
		validation.Field(&giftAssociate.VisitID, validation.Required),
		validation.Field(&giftAssociate.FriendID, validation.Required),
		validation.Field(&giftAssociate.Quantity, validation.Required),
	); err != nil {
		return "", err
	} // not empty
	return s.repo.Create(ctx, giftAssociate)
}

// Update a gift associate
func (s *Service) Update(ctx context.Context, giftAssociate types.GiftAssociate) error {
	if err := validation.ValidateStruct(&giftAssociate,
		validation.Field(&giftAssociate.ID, validation.Required),
		validation.Field(&giftAssociate.VisitID, validation.Required),
		validation.Field(&giftAssociate.FriendID, validation.Required),
		validation.Field(&giftAssociate.Quantity, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, giftAssociate)
}

// Delete a gift associate
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
