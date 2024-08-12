package service

import (
	"context"
	"time"

	"github.com/zhetkerbaevan/study-mongodb/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) *models.Todo {
	client = mongo
	return &models.Todo{}
}

type TodoService struct {
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func returnCollectionPointer() *mongo.Collection {
	collection := client.Database("study_db").Collection("todos") //Pointer to collection
	return collection
}

//mongodb://admin:1234@localhost:27017/study_db?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false

func (s *TodoService) InsertTodo(todo models.TodoPayload) error {
	collection := returnCollectionPointer()

	_, err := collection.InsertOne(context.Background(), models.Todo{
		Task:      todo.Task,
		Completed: todo.Completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) GetTodos() ([]primitive.M, error) {
	collection := returnCollectionPointer()
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var todos []primitive.M
	for cursor.Next(context.Background()) {
		var todo bson.M
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
