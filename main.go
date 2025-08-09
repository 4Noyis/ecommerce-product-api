package main

import (
	"log"

	"github.com/4Noyis/ecommerce-product-api/internal/repository"
	"github.com/4Noyis/ecommerce-product-api/internal/services"
	"github.com/4Noyis/ecommerce-product-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Connect to MongoDB
	client, err := database.ConnectMongoDB("ecommerce-api")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Close()

	// Create repositories
	//productRepo := repository.NewMongoProductRepository(client)
	reviewsRepo := repository.NewMongoReviewsRepository(client)

	// Create sample reviews for a sample product ID
	// You can replace this with an actual product ID from your products collection
	productID, _ := primitive.ObjectIDFromHex("aJZq8g6rGxoqSQ2J") // Sample ObjectID
	createSampleReviews(reviewsRepo, productID)

	// Create sample products using the dedicated function
	// createProductRepoSamples(productRepo)
}

func createSampleReviews(reviewsRepo *repository.MongoReviewRepository, productID primitive.ObjectID) {
	err := services.CreateSampleReviews(reviewsRepo, productID)
	if err != nil {
		log.Fatal("Failed to create sample reviews:", err)
	}
}

func createProductRepoSamples(productRepo *repository.MongoProductRepository) {
	err := services.CreateSampleProducts(productRepo)
	if err != nil {
		log.Fatal("Failed to create sample products:", err)
	}
}
