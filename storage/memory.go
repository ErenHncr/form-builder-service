package storage

import "github.com/erenhncr/go-api-structure/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
