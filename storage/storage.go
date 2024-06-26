package storage

import (
	"context"
	"log"
	"os"

	"github.com/erenhncr/go-api-structure/types"
)

const (
	postgres = "postgres"
	mongodb  = "mongodb"
	memory   = "memory"
)

type Storage interface {
	Connect(context.Context) error
	Disconnect(context.Context) error
	GetQuestions(types.Pagination, []types.Sorting) ([]types.Question, int, error)
	GetQuestion(string) (*types.Question, error)
	CreateQuestion(types.Question) (*types.Question, error)
	DeleteQuestion(string) error
	UpdateQuestion(string, types.QuestionPatch) (*types.Question, error)
}

func NewStorage() Storage {
	engineName := os.Getenv("DATABASE")

	var storage struct {
		name   string
		engine Storage
	}

	switch engineName {
	case postgres:
		storage.name = postgres
		storage.engine = &PostgresStorage{}
	case mongodb:
		storage.name = mongodb
		storage.engine = &MongoDBStorage{}
	default:
		storage.name = memory
		storage.engine = &MemoryStorage{}
	}

	log.Printf("database engine: %v", storage.name)

	return storage.engine
}
