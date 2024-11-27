package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEnvVar(envVar string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	return os.Getenv(envVar)
}

func ConnectToMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongoDBURI := GetEnvVar("MONGODBURI")
	fmt.Println("URI:", mongoDBURI)
	clientOptions := options.Client().ApplyURI(mongoDBURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error creating MongoDB client")
		return nil
	}
	connectionErr := client.Ping(ctx, nil)
	if connectionErr != nil {
		log.Fatal("Connection error to MongoDB", connectionErr)
		return nil
	}
	return client
}
