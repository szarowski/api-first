package user

import "sync/atomic"

type StoreUser struct {
	maxId atomic.Uint32
	db    map[uint32]any
}

func NewStoreUser(userStorage map[uint32]any) *StoreUser {
	return &StoreUser{
		db: userStorage,
	}
}

func (us *StoreUser) NewId() uint32 {
	return us.maxId.Add(1)
}
