package usercache

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
)

type (
	UsercacheReponsitory interface {
		FindUserByUserName(username string) (*types.Usercache, error)
		InserUser(user *types.Usercache) error
	}
	UserCacheService struct {
		repo   UsercacheReponsitory
		logger glog.Logger
	}
)

func NewUserCacheService(r UsercacheReponsitory, l glog.Logger) *UserCacheService {
	return &UserCacheService{
		repo:   r,
		logger: l,
	}
}
func (service *UserCacheService) GetUserByeUserName(username string) (*types.Usercache, error) {
	return service.repo.FindUserByUserName(username)
}
func (service *UserCacheService) AddUser(user *types.Usercache) error {
	return service.repo.InserUser(user)
}
