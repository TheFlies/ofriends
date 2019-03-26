package visit

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a visit repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Visit, error)
	FindAll(ctx context.Context) ([]types.Visit, error)
	Create(ctx context.Context, visit types.Visit) (string, error)
	Update(ctx context.Context, visit types.Visit) error
	Delete(ctx context.Context, id string) error
}

// Service is an visit service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new visit service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given visit by his/her id
func (s *Service) Get(ctx context.Context, id string) (*types.Visit, error) {
	return s.repo.FindByID(ctx, id)
}

// Get All return all visits from database
func (s *Service) GetAll(ctx context.Context) ([]types.Visit, error) {
	return s.repo.FindAll(ctx)
}

// Create a visit
func (s *Service) Create(ctx context.Context, visit types.Visit) (string, error) {
	if err := validation.ValidateStruct(&visit,

	); err != nil {
		return "",err
	} // not empty
	return s.repo.Create(ctx, visit)
}

// Update a visit
func (s *Service) Update(ctx context.Context, visit types.Visit) error {
	if err := validation.ValidateStruct(&visit,

	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, visit)
}

// Delete a visit
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
