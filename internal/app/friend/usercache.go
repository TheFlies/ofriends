package friend

import (
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
)

type (
	UseCacheRepository interface {
		FindUserByID(id string) (*types.Usercache, error)
		InserUser(use *types.Usercache) error
		FindUserByUserName(usename string) (*types.Usercache, error)
	}

	Userservice struct {
		repo   UseCacheRepository
		logger glog.Logger
	}
)

func NewUserService(r UseCacheRepository, l glog.Logger) *Userservice {
	return &Userservice{
		repo:   r,
		logger: l,
	}
}
