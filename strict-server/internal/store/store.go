package store

import "strict-server/internal/store/user"

type StorageType int

const (
	Users StorageType = iota
)

type Store struct {
	Storage map[StorageType]map[uint32]any
	// Only Store for Users
	User *user.StoreUser
}

func NewStore() *Store {
	// For simplicity - Storage as map in memory
	storage := map[StorageType]map[uint32]any{
		Users: make(map[uint32]any),
	}
	return &Store{
		Storage: storage,
		User:    user.NewStoreUser(storage[Users]),
	}
}
