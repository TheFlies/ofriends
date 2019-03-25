package user

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
)

type (
	UserRepository interface {
		FindUserByUserName(username string) (*types.User, error)
		InserUser(user *types.User) error
		CheckUserByUsername(username string) bool
		UpdateUser(user *types.User) error
	}
	UserService struct {
		repo   UserRepository
		logger glog.Logger
	}
)

func NewUserService(r UserRepository, l glog.Logger) *UserService {
	return &UserService{
		repo:   r,
		logger: l,
	}
}
func (s *UserService) GetByName(username string) (*types.User, error) {
	return s.repo.FindUserByUserName(username)
}
func (s *UserService) AddUser(user *types.User) error {
	return s.repo.InserUser(user)
}
func (s *UserService) CheckExistence(username string) bool {
	return s.repo.CheckUserByUsername(username)
}
func (s *UserService) ChangeUserPassword(user *types.User) error {
	return s.repo.UpdateUser(user)
}
