package gift

import (
	"context"
	"errors"

	"github.com/TheFlies/ofriends/internal/app/giftassociate"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a gift repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Gift, error)
	FindAll(ctx context.Context) ([]types.Gift, error)
	Create(ctx context.Context, gift types.Gift) (string, error)
	Update(ctx context.Context, gift types.Gift) error
	Delete(ctx context.Context, id string) error
}

// Service is an gift service
type Service struct {
	repo      Repository
	assocRepo giftassociate.Repository
	logger    glog.Logger
}

// NewService return a new gift service
func NewService(r Repository, assocRepo giftassociate.Repository, l glog.Logger) *Service {
	return &Service{
		repo:      r,
		assocRepo: assocRepo,
		logger:    l,
	}
}

// Get return given gift by his/her id
func (s *Service) Get(ctx context.Context, id string) (*types.Gift, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAll return all gifts from database
func (s *Service) GetAll(ctx context.Context) ([]types.Gift, error) {
	return s.repo.FindAll(ctx)
}

// Create a gifts
func (s *Service) Create(ctx context.Context, gift types.Gift) (string, error) {
	if err := validation.ValidateStruct(&gift,
		validation.Field(&gift.Name, validation.Required),
	); err != nil {
		return "", err
	} // not empty
	return s.repo.Create(ctx, gift)
}

// Update a gift
func (s *Service) Update(ctx context.Context, gift types.Gift) error {
	if err := validation.ValidateStruct(&gift,
		validation.Field(&gift.ID, validation.Required),
		validation.Field(&gift.Name, validation.Required),
	); err != nil {
		return err
	} // not empty
	if oldGift, err := s.repo.FindByID(ctx, gift.ID); err == nil {
		error := s.repo.Update(ctx, gift)
		if error == nil && oldGift.Name != gift.Name {
			return s.assocRepo.UpdateNameByGiftID(ctx, gift.Name, gift.ID)
		}
	}
	return nil
}

// Delete a gift
func (s *Service) Delete(ctx context.Context, id string) error {
	if !s.assocRepo.IsAssignedGift(ctx, id, "") {
		return s.repo.Delete(ctx, id)
	}
	return errors.New("This gift has assigned")
}
