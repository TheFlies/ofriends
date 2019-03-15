package user

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
)

type (
	UserReponsitory interface {
		FindUserByUserName(username string) (*types.User, error)
		InserUser(user *types.User) error
		CheckUserByUsername(username string) bool
	}
	UserService struct {
		repo   UserReponsitory
		logger glog.Logger
	}
)

func NewUserService(r UserReponsitory, l glog.Logger) *UserService {
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
