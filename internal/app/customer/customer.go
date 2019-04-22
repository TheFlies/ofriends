package customer

import (
	"context"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/go-ozzo/ozzo-validation"
)

// Repository is an interface of a customer repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Customer, error)
	FindAll(ctx context.Context) ([]types.Customer, error)
	Create(ctx context.Context, customer types.Customer) (string, error)
	Update(ctx context.Context, customer types.Customer) error
	Delete(ctx context.Context, id string) error
}

// Service is an customer service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new customer service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given customer by his/her id
func (s *Service) Get(ctx context.Context, id string) (*types.Customer, error) {
	return s.repo.FindByID(ctx, id)
}

// Get All return all customer from database
func (s *Service) GetAll(ctx context.Context) ([]types.Customer, error) {
	return s.repo.FindAll(ctx)
}

// Create a customer
func (s *Service) Create(ctx context.Context, customer types.Customer) (string, error) {
	if err := validation.ValidateStruct(&customer,
		validation.Field(&customer.Name, validation.Required),
		validation.Field(&customer.Title, validation.Required),
		validation.Field(&customer.Position, validation.Required),
		validation.Field(&customer.Project, validation.Required),
	); err != nil {
		return "",err
	} // not empty
	return s.repo.Create(ctx, customer)
}

// Update a customer
func (s *Service) Update(ctx context.Context, customer types.Customer) error {
	if err := validation.ValidateStruct(&customer,
		validation.Field(&customer.ID, validation.Required),
		validation.Field(&customer.Name, validation.Required),
		validation.Field(&customer.Title, validation.Required),
		validation.Field(&customer.Position, validation.Required),
		validation.Field(&customer.Project, validation.Required),
	); err != nil {
		return err
	} // not empty
	return s.repo.Update(ctx, customer)
}

// Delete a customer
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
