package main

import (
	"context"
	"log"
	"time"

	"github.com/zhetkerbaevan/study-mongodb/cmd/api"
	"github.com/zhetkerbaevan/study-mongodb/internal/db"
	"github.com/zhetkerbaevan/study-mongodb/internal/service"
)

func main() {
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	service.New(mongoClient)

	apiServer := api.NewAPIServer(":8080")
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
