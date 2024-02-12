package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strict-server/internal/api/gen"
	"strict-server/internal/errors"
	"testing"
)

var (
	userToGet = gen.GetUserRequestObject{
		Id: 1,
	}
	invalidUserToGet = gen.GetUserRequestObject{
		Id: 2,
	}
)

func TestControllerUser_GetUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		request     gen.GetUserRequestObject
		expected    *gen.GetUser200JSONResponse
		expectedErr error
	}{
		{
			name:    "Should get user",
			request: userToGet,
			expected: &gen.GetUser200JSONResponse{
				Id:        1,
				FirstName: "John",
				LastName:  "Doe",
				UserEmail: "john.doe@myemail.com",
			},
			expectedErr: nil,
		},
		{
			name:        "Should not get user",
			request:     invalidUserToGet,
			expected:    nil,
			expectedErr: errors.ErrStatusNotFoundUser,
		},
	}

	userController := createUser(t, userToCreate)

	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userController.GetUser(ctx, test.request)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, *test.expected, actual)
		})
	}
}
