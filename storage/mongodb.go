package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/erenhncr/go-api-structure/types"
	"github.com/erenhncr/go-api-structure/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	questionCollection = "questions"
)

type MongoDBStorage struct {
	client  *mongo.Client
	context context.Context
}

func (s *MongoDBStorage) getCollection(collectionName string) *mongo.Collection {
	return s.client.Database(util.GetDatabaseName()).Collection(collectionName)
}

func (s *MongoDBStorage) Connect(ctx context.Context) error {
	databaseUrl := util.GetDatabaseURL()
	databaseName := util.GetDatabaseName()

	if databaseName == "" {
		return fmt.Errorf("database name cannot be empty")
	}
	if databaseUrl == "" {
		return fmt.Errorf("database url cannot be empty")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseUrl))
	if err != nil {
		return fmt.Errorf("database connection error: %v", err)
	}

	var result bson.M
	if err := client.Database(databaseName).RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return fmt.Errorf("database ping error: %v", err)
	}

	s.context = ctx
	s.client = client
	log.Println("pinged your deployment. successfully connected to mongodb")

	return nil
}

func (s *MongoDBStorage) Disconnect(ctx context.Context) error {
	if err := s.client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func (s *MongoDBStorage) GetQuestions(pagination types.Pagination) []types.Question {
	return []types.Question{}
}

func (s *MongoDBStorage) GetQuestion(id string) (*types.Question, error) {
	return nil, nil
}

func (s *MongoDBStorage) CreateQuestion(question types.Question) (*types.Question, error) {
	collection := s.getCollection(questionCollection)

	result, err := collection.InsertOne(s.context, question)
	if err != nil {
		return nil, fmt.Errorf("create_question_error %v", err.Error())
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	question.ID = insertedId

	return &question, nil
}
func (s *MongoDBStorage) UpdateQuestion(id string, q types.Question) (*types.Question, error) {
	return nil, nil
}

func (s *MongoDBStorage) DeleteQuestion(id string) error {
	return nil
}
