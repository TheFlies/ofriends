package activity

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a activity repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Activity, error)
	FindAll(ctx context.Context) ([]types.Activity, error)
	FindByVisitID(ctx context.Context, visitId string) ([]types.Activity, error)
	Create(ctx context.Context, act types.Activity) (string, error)
	Update(ctx context.Context, act types.Activity) error
	Delete(ctx context.Context, id string) error
}

// Service is an activity service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new activity service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given activity by activity id
func (s *Service) Get(ctx context.Context, id string) (*types.Activity, error) {
	return s.repo.FindByID(ctx, id)
}

// Get all activities from database by visit ID
func (s *Service) GetByVisitID(ctx context.Context, visitId string) ([]types.Activity, error) {
	return s.repo.FindByVisitID(ctx, visitId)
}

// Get All return all activities from database
func (s *Service) GetAll(ctx context.Context) ([]types.Activity, error) {
	return s.repo.FindAll(ctx)
}

// Create a activity
func (s *Service) Create(ctx context.Context, act types.Activity) (string, error) {
	if err := validation.ValidateStruct(&act,
		validation.Field(&act.StartDate, validation.Required),
		validation.Field(&act.EndDate, validation.Required),
		validation.Field(&act.VisitID, validation.Required),
	); err != nil {
		return "",err
	} // not empty
	return s.repo.Create(ctx, act)
}

// Update a activity
func (s *Service) Update(ctx context.Context, act types.Activity) error {
	if err := validation.ValidateStruct(&act,
		validation.Field(&act.ID, validation.Required),
		validation.Field(&act.StartDate, validation.Required),
		validation.Field(&act.EndDate, validation.Required),
		validation.Field(&act.VisitID, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, act)
}

// Delete a activity
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
