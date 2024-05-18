package storage

import "github.com/erenhncr/go-api-structure/types"

type MongoStorage struct{}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
