package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strict-server/internal/api/gen"
	"strict-server/internal/store"
	"testing"
)

var (
	user1 = gen.CreateUserRequestObject{
		Body: &gen.User{
			FirstName: "John",
			LastName:  "Doe",
			UserEmail: "john.doe@myemail.com",
		},
	}
)

func TestControllerUser_CreateUser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		request     gen.CreateUserRequestObject
		expected    *gen.CreateUser201JSONResponse
		expectedErr error
	}{
		{
			name:    "Should create user1",
			request: user1,
			expected: &gen.CreateUser201JSONResponse{
				Id:        1,
				FirstName: "John",
				LastName:  "Doe",
				UserEmail: "john.doe@myemail.com",
			},
			expectedErr: nil,
		},
	}

	userStore := store.NewStore()
	userController := NewControllerUser(userStore.User)

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userController.CreateUser(ctx, test.request)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, *test.expected, actual)
		})
	}
}
