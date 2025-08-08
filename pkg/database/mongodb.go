package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func ConnectMongoDB(dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	uri := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	return &MongoDB{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (db *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}

// Collections
func (db *MongoDB) ProductsCollection() *mongo.Collection {
	return db.Database.Collection("products")
}

func (db *MongoDB) ReviewsCollection() *mongo.Collection {
	return db.Database.Collection("reviews")
}

func (db *MongoDB) CartCollection() *mongo.Collection {
	return db.Database.Collection("cart")
}

func (db *MongoDB) UsersCollection() *mongo.Collection {
	return db.Database.Collection("users")
}
