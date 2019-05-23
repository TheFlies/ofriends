package customer

import (
	"context"
	"errors"

	"github.com/TheFlies/ofriends/internal/app/cusvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	validation "github.com/go-ozzo/ozzo-validation"
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
	repo      Repository
	assocRepo cusvisitassoc.Repository
	logger    glog.Logger
}

// NewService return a new customer service
func NewService(r Repository, assocRepo cusvisitassoc.Repository, l glog.Logger) *Service {
	return &Service{
		repo:      r,
		assocRepo: assocRepo,
		logger:    l,
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
		return "", err
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

	if cus, err := s.repo.FindByID(ctx, customer.ID); err == nil {
		error := s.repo.Update(ctx, customer)
		if error == nil && cus.Name != customer.Name {
			return s.assocRepo.UpdateNameByCusID(ctx, customer.Name, customer.ID)
		}
	}
	return nil
}

// Delete a customer
func (s *Service) Delete(ctx context.Context, id string) error {
	if !s.assocRepo.IsAssignedCustomer(ctx, id, "") {
		return s.repo.Delete(ctx, id)
	}
	return errors.New("This customer has assigned with visit")
}
