package storage

import (
	"context"
	"fmt"
	"log"
	"time"

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

func getObjectID(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("invalid_id")
	}
	return objectId, nil
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

	opts := options.Client()
	opts.SetConnectTimeout(30 * time.Second)
	client, err := mongo.Connect(ctx, opts.ApplyURI(databaseUrl))
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

func (s *MongoDBStorage) GetQuestions(pagination types.Pagination) ([]types.Question, int, error) {
	collection := s.getCollection(questionCollection)
	filter := bson.D{}

	totalItems, err := collection.CountDocuments(s.context, filter)
	if err != nil {
		return nil, 0, fmt.Errorf(types.ErrorCode[types.ErrorCodeInternalServerError])
	}

	itemCursor := (pagination.Page * pagination.Size) - pagination.Size
	if int(totalItems)-itemCursor <= 0 {
		return []types.Question{}, int(totalItems), nil
	}

	opts := options.Find()
	opts.SetSort(bson.M{"createdat": -1})
	opts.SetLimit(int64(pagination.Size)).SetSkip(int64(pagination.Page - 1))
	opts.SetMaxAwaitTime(30 * time.Second)

	var results []types.Question
	cursor, err := collection.Find(s.context, filter, opts)
	if err != nil {
		return nil, 0, fmt.Errorf(types.ErrorCode[types.ErrorCodeInternalServerError])
	}

	if err = cursor.All(s.context, &results); err != nil {
		return nil, 0, fmt.Errorf(types.ErrorCode[types.ErrorCodeInternalServerError])
	}

	return results, int(totalItems), nil
}

func (s *MongoDBStorage) GetQuestion(id string) (*types.Question, error) {
	objectId, err := getObjectID(id)

	if err != nil {
		return nil, err
	}

	collection := s.getCollection(questionCollection)
	filter := bson.M{"_id": objectId}

	var question *types.Question
	err = collection.FindOne(s.context, filter).Decode(&question)

	if err != nil {
		return nil, fmt.Errorf("not_found")
	}

	return question, nil
}

func (s *MongoDBStorage) CreateQuestion(question types.Question) (*types.Question, error) {
	collection := s.getCollection(questionCollection)

	ctx, cancel := context.WithTimeout(s.context, 15*time.Second)
	defer cancel()

	createdAt := time.Now()
	question.CreatedAt = createdAt
	question.UpdatedAt = createdAt

	result, err := collection.InsertOne(ctx, question)
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
	objectId, err := getObjectID(id)

	if err != nil {
		return err
	}

	collection := s.getCollection(questionCollection)
	filter := bson.M{"_id": objectId}
	_, err = collection.DeleteOne(s.context, filter)

	if err != nil {
		return err
	}

	return nil
}
