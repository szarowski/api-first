package user

import (
	"strict-server/internal/store/user"
)

type ControllerUser struct {
	storeUser *user.StoreUser
}

func NewControllerUser(userStore *user.StoreUser) *ControllerUser {
	return &ControllerUser{
		storeUser: userStore,
	}
}
