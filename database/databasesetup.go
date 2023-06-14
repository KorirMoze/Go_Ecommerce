package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err !=nil {
		log.Fatal(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()
}

func UserData(client *mongo.client,collectionName string) *mongo.Collection{

}
func ProductData(client *mongo.client,collectionNAme string) *mongo.Collection{

}