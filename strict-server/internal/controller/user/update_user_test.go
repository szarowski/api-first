package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strict-server/internal/api/gen"
	"strict-server/internal/errors"
	"strict-server/internal/store"
	"testing"
)

var (
	userToCreate = gen.CreateUserJSONRequestBody{
		FirstName: "John",
		LastName:  "Doe",
		UserEmail: "john.doe@myemail.com",
	}
	userToUpdate = gen.UpdateUserRequestObject{
		Id: 1,
		Body: &gen.User{
			FirstName: "Jack",
			LastName:  "Poe",
			UserEmail: "jack.poe@myemail.com",
		},
	}
	invalidUser = gen.UpdateUserRequestObject{
		Id: 2,
		Body: &gen.User{
			FirstName: "Jack",
			LastName:  "Poe",
			UserEmail: "jack.poe@myemail.com",
		},
	}
)

func TestControllerUser_UpdateUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		request     gen.UpdateUserRequestObject
		expected    *gen.UpdateUser200JSONResponse
		expectedErr error
	}{
		{
			name:    "Should update user",
			request: userToUpdate,
			expected: &gen.UpdateUser200JSONResponse{
				Id:        1,
				FirstName: "Jack",
				LastName:  "Poe",
				UserEmail: "jack.poe@myemail.com",
			},
			expectedErr: nil,
		},
		{
			name:        "Should not update user",
			request:     invalidUser,
			expected:    nil,
			expectedErr: errors.ErrStatusNotFoundUser,
		},
	}

	userController := createUser(t, userToCreate)

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userController.UpdateUser(ctx, test.request)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, *test.expected, actual)
		})
	}
}

func createUser(t *testing.T, user gen.CreateUserJSONRequestBody) *ControllerUser {
	userStore := store.NewStore()
	userController := NewControllerUser(userStore.User)
	_, err := userController.storeUser.CreateUser(user)
	if err != nil {
		t.Fatalf("User.CreateUser() error = %v", err)
	}
	return userController
}
