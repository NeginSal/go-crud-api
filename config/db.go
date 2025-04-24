package config

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

var DB *mongo.Database

func ConnectDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Get URI and database name from .env file
	mongoURI := os.Getenv("MONGO_URI")
	databaseName := os.Getenv("DB_NAME")

	if mongoURI == "" {
		log.Fatal("❌ MONGO_URI is not set in .env")
	}
	if databaseName == "" {
		log.Fatal("❌ DB_NAME is not set in .env")
	}

	// Set up connection options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	// Ping the database to test connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB ping error:", err)
	}

	fmt.Println("✅ Connected to MongoDB Atlas!")

	// Set the selected database
	DB = client.Database(databaseName)
}
