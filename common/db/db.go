package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arthuruan/training-consultancy/common/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	uri := configs.EnvMongoURI()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client
}

// Getting Database Collections
func GetCollection(client *mongo.Client, collecitonName string) *mongo.Collection {
	return client.Database("TrainingConsultancy").Collection(collecitonName)
}
