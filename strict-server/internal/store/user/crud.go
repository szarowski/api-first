package user

import (
	"strict-server/internal/api/gen"
	"strict-server/internal/errors"
)

func (us *StoreUser) CreateUser(user gen.User) (*gen.UserWithId, error) {
	id := us.NewId()
	if us.db[id] == nil {
		userWithId := gen.UserWithId{
			Id:        id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserEmail: user.UserEmail,
		}
		us.db[id] = userWithId
		return &userWithId, nil
	}
	return nil, errors.ErrStatusConflictUser
}

func (us *StoreUser) UpdateUser(id uint32, user gen.User) (*gen.UserWithId, error) {
	if us.db[id] != nil {
		userWithId := gen.UserWithId{
			Id:        id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserEmail: user.UserEmail,
		}
		us.db[id] = userWithId
		return &userWithId, nil
	}
	return nil, errors.ErrStatusNotFoundUser
}

func (us *StoreUser) GetUser(id uint32) (*gen.UserWithId, error) {
	if us.db[id] != nil {
		userWithId := us.db[id].(gen.UserWithId)
		return &userWithId, nil
	}
	return nil, errors.ErrStatusNotFoundUser
}

func (us *StoreUser) DeleteUser(id uint32) error {
	if us.db[id] != nil {
		us.db[id] = nil
		return nil
	}
	return errors.ErrStatusNotFoundUser
}
