package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strict-server/internal/api/gen"
	"strict-server/internal/errors"
	"testing"
)

var (
	userToDelete = gen.DeleteUserRequestObject{
		Id: 1,
	}
	invalidUserToDelete = gen.DeleteUserRequestObject{
		Id: 2,
	}
)

func TestControllerUser_DeleteUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		request     gen.DeleteUserRequestObject
		expected    *gen.DeleteUser204Response
		expectedErr error
	}{
		{
			name:        "Should delete user",
			request:     userToDelete,
			expected:    &gen.DeleteUser204Response{},
			expectedErr: nil,
		},
		{
			name:        "Should not delete user",
			request:     invalidUserToDelete,
			expected:    nil,
			expectedErr: errors.ErrStatusNotFoundUser,
		},
	}

	userController := createUser(t, userToCreate)

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userController.DeleteUser(ctx, test.request)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, *test.expected, actual)
		})
	}
}
