package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connString = "mongodb://admin:1234@localhost:27017/study_db?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false"
const dbName = "study_db"
const colName = "todos"

var Collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connString) //MongoDB connection string

	client, err := mongo.Connect(context.TODO(), clientOptions) //Connect to MongoDB
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, err
}
