package actvisitassoc

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a activity visit repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.ActVisitAssoc, error)
	FindAll(ctx context.Context) ([]types.ActVisitAssoc, error)
	FindByVisitID(ctx context.Context, visitID string) ([]types.ActVisitAssoc, error)
	FindByActID(ctx context.Context, actID string) ([]types.ActVisitAssoc, error)
	Create(ctx context.Context, actVisitAssoc types.ActVisitAssoc) (string, error)
	Update(ctx context.Context, actVisitAssoc types.ActVisitAssoc) error
	Delete(ctx context.Context, id string) error
	DeleteByVisitID(ctx context.Context, visitID string) error
}

// Service is an activity visit associate service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new activity visit associate service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given activity associate by id
func (s *Service) Get(ctx context.Context, id string) (*types.ActVisitAssoc, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAll return all activity visit associates from database
func (s *Service) GetAll(ctx context.Context) ([]types.ActVisitAssoc, error) {
	return s.repo.FindAll(ctx)
}

// GetByVisitID return given activity associate by visit id
func (s *Service) GetByVisitID(ctx context.Context, visitID string) ([]types.ActVisitAssoc, error) {
	return s.repo.FindByVisitID(ctx, visitID)
}

// GetByActID return given activity associate by activity id
func (s *Service) GetByActID(ctx context.Context, actID string) ([]types.ActVisitAssoc, error) {
	return s.repo.FindByActID(ctx, actID)
}

// Create a activity visit associates
func (s *Service) Create(ctx context.Context, actVisitAssoc types.ActVisitAssoc) (string, error) {
	if err := validation.ValidateStruct(&actVisitAssoc,
		validation.Field(&actVisitAssoc.VisitID, validation.Required),
	); err != nil {
		return "", err
	} // not empty
	return s.repo.Create(ctx, actVisitAssoc)
}

// Update a activity visit associate
func (s *Service) Update(ctx context.Context, actVisitAssoc types.ActVisitAssoc) error {
	if err := validation.ValidateStruct(&actVisitAssoc,
		validation.Field(&actVisitAssoc.ID, validation.Required),
		validation.Field(&actVisitAssoc.VisitID, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, actVisitAssoc)
}

// Delete a activity visit associate
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// Delete a activity visit associate
func (s *Service) DeleteByVisitID(ctx context.Context, visitId string) error {
	return s.repo.DeleteByVisitID(ctx, visitId)
}
