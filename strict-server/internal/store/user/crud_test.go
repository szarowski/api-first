package user

import (
	"github.com/stretchr/testify/assert"
	"strict-server/internal/api/gen"
	"strict-server/internal/errors"
	"testing"
)

var (
	user1 = gen.User{
		FirstName: "John",
		LastName:  "Doe",
		UserEmail: "john.doe@myemail.com",
	}
	user2 = gen.User{
		FirstName: "Jack",
		LastName:  "Poe",
		UserEmail: "jack.poe@myemail.com",
	}
)

func TestStoreUser_CreateUser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		request     gen.User
		expected    *gen.UserWithId
		expectedErr error
	}{
		{
			name:    "Should create user1",
			request: user1,
			expected: &gen.UserWithId{
				Id:        1,
				FirstName: "John",
				LastName:  "Doe",
				UserEmail: "john.doe@myemail.com",
			},
			expectedErr: nil,
		},
	}

	userStore := NewStoreUser(map[uint32]any{})

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userStore.CreateUser(test.request)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestStoreUser_UpdateUser(t *testing.T) {
	t.Parallel()
	type args struct {
		id uint32
	}

	tests := []struct {
		name        string
		args        args
		request     gen.User
		expected    *gen.UserWithId
		expectedErr error
	}{
		{
			name: "Should update user1",
			args: args{
				id: 1,
			},
			request: user2,
			expected: &gen.UserWithId{
				Id:        1,
				FirstName: "Jack",
				LastName:  "Poe",
				UserEmail: "jack.poe@myemail.com",
			},
			expectedErr: nil,
		},
		{
			name: "Should not update user1",
			args: args{
				id: 2,
			},
			request:     user2,
			expected:    nil,
			expectedErr: errors.ErrStatusNotFoundUser,
		},
	}

	userStore := createUser(t, user1)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userStore.UpdateUser(test.args.id, test.request)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestStoreUser_GetUser(t *testing.T) {
	t.Parallel()
	type args struct {
		id uint32
	}

	tests := []struct {
		name        string
		args        args
		expected    *gen.UserWithId
		expectedErr error
	}{
		{
			name: "Should get user1",
			args: args{
				id: 1,
			},
			expected: &gen.UserWithId{
				Id:        1,
				FirstName: "John",
				LastName:  "Doe",
				UserEmail: "john.doe@myemail.com",
			},
			expectedErr: nil,
		},
		{
			name: "Should not get user1",
			args: args{
				id: 2,
			},
			expected:    nil,
			expectedErr: errors.ErrStatusNotFoundUser,
		},
	}

	userStore := createUser(t, user1)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := userStore.GetUser(test.args.id)
			if err != nil {
				assert.Equal(t, test.expectedErr, err)
				return
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestStoreUser_DeleteUser(t *testing.T) {
	t.Parallel()
	type args struct {
		id uint32
	}

	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "Should delete user1",
			args: args{
				id: 1,
			},
			expectedErr: nil,
		},
		{
			name: "Should not delete user1",
			args: args{
				id: 2,
			},
			expectedErr: errors.ErrStatusNotFoundUser,
		},
	}

	userStore := createUser(t, user1)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := userStore.DeleteUser(test.args.id)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func createUser(t *testing.T, user gen.User) *StoreUser {
	userStore := NewStoreUser(map[uint32]any{})
	_, err := userStore.CreateUser(user)
	if err != nil {
		t.Fatalf("User.CreateUser() error = %v", err)
	}
	return userStore
}
