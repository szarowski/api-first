package user

import (
	"context"
	"strict-server/internal/api/gen"
)

// CreateUser - POST /v1/users
func (uc *ControllerUser) CreateUser(_ context.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {
	result, err := uc.storeUser.CreateUser(*request.Body)
	if err != nil {
		return nil, err
	}
	return gen.CreateUser201JSONResponse(*result), err
}
