package activity

import (
	"context"
	"errors"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
)

// Repository is an interface of a activity repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Activity, error)
	FindAll(ctx context.Context) ([]types.Activity, error)
	Create(ctx context.Context, act types.Activity) (string, error)
	Update(ctx context.Context, act types.Activity) error
	Delete(ctx context.Context, id string) error
}

// Service is an activity service
type Service struct {
	repo      Repository
	assocService ActVisitAssocService
	logger    glog.Logger
}

type ActVisitAssocService interface {
	UpdateNameByActID(ctx context.Context, activityName string, activityID string) error
	IsAssignedActivity(ctx context.Context, activityID string, visitID string) bool
}

// NewService return a new activity service
func NewService(r Repository, assocService ActVisitAssocService, l glog.Logger) *Service {
	return &Service{
		repo:      r,
		assocService: assocService,
		logger:    l,
	}
}

// Get return given activity by activity id
func (s *Service) Get(ctx context.Context, id string) (*types.Activity, error) {
	return s.repo.FindByID(ctx, id)
}

// Get All return all activities from database
func (s *Service) GetAll(ctx context.Context) ([]types.Activity, error) {
	return s.repo.FindAll(ctx)
}

// Create a activity
func (s *Service) Create(ctx context.Context, act types.Activity) (string, error) {
	return s.repo.Create(ctx, act)
}

// Update a activity
func (s *Service) Update(ctx context.Context, act types.Activity) error {
	activity, err := s.repo.FindByID(ctx, act.ID)
	if err != nil {
		return err
	}
	
	if err := s.repo.Update(ctx, act); err != nil {
		return err
	}

	if activity.Name != act.Name {
		return s.assocService.UpdateNameByActID(ctx, act.Name, act.ID)
	}

	return nil
}

// Delete a activity
func (s *Service) Delete(ctx context.Context, id string) error {
	if !s.assocService.IsAssignedActivity(ctx, id, "") {
		return s.repo.Delete(ctx, id)
	}
	return errors.New("This Activity is assigned in another Visit. You cannot delete it.")

}
