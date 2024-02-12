package user

import (
	"context"
	"strict-server/internal/api/gen"
)

// UpdateUser - PUT /v1/users/{id}
func (uc *ControllerUser) UpdateUser(_ context.Context, request gen.UpdateUserRequestObject) (gen.UpdateUserResponseObject, error) {
	result, err := uc.storeUser.UpdateUser(request.Id, *request.Body)
	if err != nil {
		return nil, err
	}
	return gen.UpdateUser200JSONResponse(*result), err
}
