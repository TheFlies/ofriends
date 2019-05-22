package cusvisitassoc

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a customer visit repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.CusVisitAssoc, error)
	FindAll(ctx context.Context) ([]types.CusVisitAssoc, error)
	FindByVisitID(ctx context.Context, visitID string) ([]types.CusVisitAssoc, error)
	Create(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) (string, error)
	Update(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) error
	Delete(ctx context.Context, id string) error
	DeleteByVisitID(ctx context.Context, visitID string) error
}

// Service is an customer visit associate service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new customer visit associate service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given gift associate by id
func (s *Service) Get(ctx context.Context, id string) (*types.CusVisitAssoc, error) {
	return s.repo.FindByID(ctx, id)
}

// GetAll return all customer visit associates from database
func (s *Service) GetAll(ctx context.Context) ([]types.CusVisitAssoc, error) {
	return s.repo.FindAll(ctx)
}

// GetByVisitID return given gift associate by visit id
func (s *Service) GetByVisitID(ctx context.Context, visitID string) ([]types.CusVisitAssoc, error) {
	return s.repo.FindByVisitID(ctx, visitID)
}

// Create a customer visit associates
func (s *Service) Create(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) (string, error) {
	if err := validation.ValidateStruct(&cusVisitAssoc,
		validation.Field(&cusVisitAssoc.VisitID, validation.Required),
	); err != nil {
		return "", err
	} // not empty
	return s.repo.Create(ctx, cusVisitAssoc)
}

// Update a customer visit associate
func (s *Service) Update(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) error {
	if err := validation.ValidateStruct(&cusVisitAssoc,
		validation.Field(&cusVisitAssoc.ID, validation.Required),
		validation.Field(&cusVisitAssoc.VisitID, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, cusVisitAssoc)
}

// Delete a customer visit associate
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// Delete a customer visit associate
func (s *Service) DeleteByVisitID(ctx context.Context, visitId string) error {
	return s.repo.DeleteByVisitID(ctx, visitId)
}
