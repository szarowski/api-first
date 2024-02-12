package user

import (
	"context"
	"strict-server/internal/api/gen"
)

// GetUser - GET /v1/users/{id}
func (uc *ControllerUser) GetUser(_ context.Context, request gen.GetUserRequestObject) (gen.GetUserResponseObject, error) {
	result, err := uc.storeUser.GetUser(request.Id)
	if err != nil {
		return nil, err
	}
	return gen.GetUser200JSONResponse(*result), nil
}
