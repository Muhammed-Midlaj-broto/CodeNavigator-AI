package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // fallback for development
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// Ping to verify
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}

	fmt.Println("Connected to MongoDB!")

	Client = client
	DB = client.Database("codebase_analyzer") // Use an appropriate database name
}
