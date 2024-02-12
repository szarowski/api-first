package controller

import (
	"strict-server/internal/api/gen"
	"strict-server/internal/controller/user"
	"strict-server/internal/store"
)

var _ gen.StrictServerInterface = (*StrictServerController)(nil)

type StrictServerController struct {
	*user.ControllerUser
}

func NewStrictServerController(store *store.Store) *StrictServerController {
	return &StrictServerController{
		user.NewControllerUser(store.User),
	}
}
