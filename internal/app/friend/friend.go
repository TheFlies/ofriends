package friend

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a friend repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Friend, error)
	FindAll(ctx context.Context) ([]types.Friend, error)
	Create(ctx context.Context, friend types.Friend) (string, error)
	Update(ctx context.Context, friend types.Friend) error
	Delete(ctx context.Context, id string) error
}

// Service is an friend service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new friend service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given friend by his/her id
func (s *Service) Get(ctx context.Context, id string) (*types.Friend, error) {
	return s.repo.FindByID(ctx, id)
}

// Get All return all friends from database
func (s *Service) GetAll(ctx context.Context) ([]types.Friend, error) {
	return s.repo.FindAll(ctx)
}

// Create a friend
func (s *Service) Create(ctx context.Context, friend types.Friend) (string, error) {
	if err := validation.ValidateStruct(&friend,
		validation.Field(&friend.Name, validation.Required),
		validation.Field(&friend.Title, validation.Required),
		validation.Field(&friend.Position, validation.Required),
		validation.Field(&friend.Project, validation.Required),
	); err != nil {
		return "",err
	} // not empty
	return s.repo.Create(ctx, friend)
}

// Update a friend
func (s *Service) Update(ctx context.Context, friend types.Friend) error {
	if err := validation.ValidateStruct(&friend,
		validation.Field(&friend.Name, validation.Required),
		validation.Field(&friend.Title, validation.Required),
		validation.Field(&friend.Position, validation.Required),
		validation.Field(&friend.Project, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, friend)
}

// Delete a friend
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
