package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool               `json:"completed" bson:"completed"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type TodoPayload struct {
	Task      string `json:"task" bson:"task"`
	Completed bool   `json:"completed" bson:"completed"`
}

type Application struct {
	Todo Todo
}

type TodoServiceInterface interface {
	InsertTodo(TodoPayload) error
	GetTodos() ([]primitive.M, error)
}
