package user

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserRepository interface {
		FindUserByUserName(username string) (*types.User, error)
		FindAll() ([]types.User, error)
		Delete(id string) error
		InsertUser(user *types.User) (string, error)
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
// Get All return all user from database
func (s *UserService) GetAll() ([]types.User, error) {
	return s.repo.FindAll()
}
// Delete a user
func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
func (s *UserService) GetByName(username string) (*types.User, error) {
	return s.repo.FindUserByUserName(username)
}
func (s *UserService) AddUser(u *types.User) (string, error) {
	if u.Password == "" {
		return s.repo.InsertUser(u)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		s.logger.Errorf("hash password fail err: %v", err)
		return "", err
	}
	u.Password = string(hashedPassword)
	return s.repo.InsertUser(u)
}
func (s *UserService) CheckExistence(username string) bool {
	return s.repo.CheckUserByUsername(username)
}
func (s *UserService) ChangePassword(username string, oldPassword string, newPassword string) error {
	dbUser, err := s.repo.FindUserByUserName(username)
	if err != nil {
		s.logger.Errorf("user %v not found form database err:%v", username, err)
		return errors.Wrap(err, "user is not exits")
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(oldPassword))
	if err != nil {
		s.logger.Errorf("your password is wrong err: %v", err)
		return errors.Wrap(err, "wrong password")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		s.logger.Errorf("hash password fail err: %v", err)
		return err
	}
	dbUser.Password = string(hashedPassword)
	return s.repo.UpdateUser(dbUser)
}
func (s *UserService) SetNewPassword(username string, password string) error {
	ok := s.repo.CheckUserByUsername(username)
	if !ok {
		s.logger.Errorf("user not found form database")
		return errors.Errorf("user is not exits")
	}
	dbUser, err := s.repo.FindUserByUserName(username)
	if err != nil {
		s.logger.Errorf("hash password fail err: %v", err)
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		s.logger.Errorf("hash password fail err: %v", err)
		return err
	}
	dbUser.Password = string(hashedPassword)
	return s.repo.UpdateUser(dbUser)
}
func (s *UserService) UpdateUser(u *types.User) error {
	ok := s.repo.CheckUserByUsername(u.Username)
	if !ok {
		s.logger.Errorf("u not found form database")
		return errors.New("u is not exits")
	}
	return s.repo.UpdateUser(u)
}
