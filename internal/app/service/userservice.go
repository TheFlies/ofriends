package service

import "github.com/TheFlies/ofriends/internal/app/types"

type (
	Userservice interface {
		GetByName(username string) (*types.User, error)
		AddUser(user *types.User) error
		CheckExistence(username string) bool
	}
)
