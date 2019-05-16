package activityvisitassociate

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a gift repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.ActivityVisitAssociate, error)
	FindAll(ctx context.Context) ([]types.ActivityVisitAssociate, error)
	FindByVisitID(ctx context.Context, visitID string) ([]types.ActivityVisitAssociate, error)
	Create(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) (string, error)
	Update(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) error
	Delete(ctx context.Context, id string) error
	DeleteByVisitID(ctx context.Context, visitID string) error
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
func (s *Service) Get(ctx context.Context, id string) (*types.ActivityVisitAssociate, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAll return all gift associates from database
func (s *Service) GetAll(ctx context.Context) ([]types.ActivityVisitAssociate, error) {
	return s.repo.FindAll(ctx)
}

// GetByVisitID return given gift associate by visit id
func (s *Service) GetByVisitID(ctx context.Context, visitID string) ([]types.ActivityVisitAssociate, error) {
	return s.repo.FindByVisitID(ctx, visitID)
}

// Create a gift associates
func (s *Service) Create(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) (string, error) {
	if err := validation.ValidateStruct(&activityVisitAssociate,
		validation.Field(&activityVisitAssociate.VisitID, validation.Required),
	); err != nil {
		return "", err
	} // not empty
	return s.repo.Create(ctx, activityVisitAssociate)
}

// Update a gift associate
func (s *Service) Update(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) error {
	if err := validation.ValidateStruct(&activityVisitAssociate,
		validation.Field(&activityVisitAssociate.ID, validation.Required),
		validation.Field(&activityVisitAssociate.VisitID, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, activityVisitAssociate)
}

// Delete a gift associate
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// Delete a gift associate
func (s *Service) DeleteByVisitID(ctx context.Context, visitId string) error {
	return s.repo.DeleteByVisitID(ctx, visitId)
}
