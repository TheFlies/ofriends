package visit

import (
	"context"
	"errors"

	"github.com/TheFlies/ofriends/internal/app/actvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/cusvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a visit repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Visit, error)
	FindAll(ctx context.Context) ([]types.Visit, error)
	FindByCustomerID(ctx context.Context, customerId string) ([]types.Visit, error)
	Create(ctx context.Context, visit types.Visit) (string, error)
	Update(ctx context.Context, visit types.Visit) error
	Delete(ctx context.Context, id string) error
}

// Service is an visit service
type Service struct {
	repo         Repository
	assocCusRepo cusvisitassoc.Repository
	assocActRepo actvisitassoc.Repository
	logger       glog.Logger
}

// NewService return a new visit service
func NewService(r Repository, assocCusRepo cusvisitassoc.Repository, assocActRepo actvisitassoc.Repository, l glog.Logger) *Service {
	return &Service{
		repo:         r,
		assocCusRepo: assocCusRepo,
		assocActRepo: assocActRepo,
		logger:       l,
	}
}

// Get return given visit by his/her id
func (s *Service) Get(ctx context.Context, id string) (*types.Visit, error) {
	return s.repo.FindByID(ctx, id)
}

// Get all visits from database by customer ID
func (s *Service) GetByCustomerID(ctx context.Context, customerId string) ([]types.Visit, error) {
	return s.repo.FindByCustomerID(ctx, customerId)
}

// Get All return all visits from database
func (s *Service) GetAll(ctx context.Context) ([]types.Visit, error) {
	return s.repo.FindAll(ctx)
}

// Create a visit
func (s *Service) Create(ctx context.Context, visit types.Visit) (string, error) {
	if err := validation.ValidateStruct(&visit,
		validation.Field(&visit.Lab, validation.Required),
		validation.Field(&visit.ArrivedTime, validation.Required),
		validation.Field(&visit.DepartureTime, validation.Required),
	); err != nil {
		return "", err
	} // not empty
	return s.repo.Create(ctx, visit)
}

// Update a visit
func (s *Service) Update(ctx context.Context, visit types.Visit) error {
	if err := validation.ValidateStruct(&visit,
		validation.Field(&visit.ID, validation.Required),
		validation.Field(&visit.Lab, validation.Required),
		validation.Field(&visit.ArrivedTime, validation.Required),
		validation.Field(&visit.DepartureTime, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, visit)
}

// Delete a visit
func (s *Service) Delete(ctx context.Context, id string) error {
	if !s.assocCusRepo.IsAssignedCustomer(ctx, "", id) && !s.assocActRepo.IsAssignedActivity(ctx, "", id) {
		return s.repo.Delete(ctx, id)
	}
	return errors.New("This customer has assigned with visit")
}
