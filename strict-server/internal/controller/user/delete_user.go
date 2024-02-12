package user

import (
	"context"
	"strict-server/internal/api/gen"
)

// DeleteUser - DELETE /v1/users/{id}
func (uc *ControllerUser) DeleteUser(_ context.Context, request gen.DeleteUserRequestObject) (gen.DeleteUserResponseObject, error) {
	err := uc.storeUser.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}
	return gen.DeleteUser204Response{}, err
}
