package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/erenhncr/go-api-structure/types"
	"github.com/erenhncr/go-api-structure/util"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	client  *sql.DB
	context context.Context
}

func (s *PostgresStorage) Connect(ctx context.Context) error {
	databaseUrl := util.GetDatabaseURL()

	if databaseUrl == "" {
		return fmt.Errorf("database url cannot be empty")
	}

	client, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return fmt.Errorf("database connection error: %v", err)
	}

	if err = client.Ping(); err != nil {
		return fmt.Errorf("database ping error: %v", err)
	}

	s.context = ctx
	s.client = client
	log.Println("pinged your deployment. successfully connected to postgres")

	return nil
}

func (s *PostgresStorage) Disconnect(ctx context.Context) error {
	if err := s.client.Close(); err != nil {
		return err
	}

	return nil
}

// CreateQuestion implements Storage.
func (s *PostgresStorage) CreateQuestion(types.Question) (*types.Question, error) {
	panic("unimplemented")
}

// DeleteQuestion implements Storage.
func (s *PostgresStorage) DeleteQuestion(string) error {
	panic("unimplemented")
}

// GetQuestion implements Storage.
func (s *PostgresStorage) GetQuestion(string) (*types.Question, error) {
	panic("unimplemented")
}

// GetQuestions implements Storage.
func (s *PostgresStorage) GetQuestions(types.Pagination, []types.Sorting) ([]types.Question, int, error) {
	panic("unimplemented")
}

// UpdateQuestion implements Storage.
func (s *PostgresStorage) UpdateQuestion(string, types.QuestionPatch) (*types.Question, error) {
	panic("unimplemented")
}
