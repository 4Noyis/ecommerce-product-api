package main

import (
	"log"

	"github.com/4Noyis/ecommerce-product-api/internal/repository"
	"github.com/4Noyis/ecommerce-product-api/internal/services"
	"github.com/4Noyis/ecommerce-product-api/pkg/database"
)

func main() {
	// Connect to MongoDB
	client, err := database.ConnectMongoDB("ecommerce-api")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Close()

	// Create repository
	productRepo := repository.NewMongoProductRepository(client)

	// Create sample products using the dedicated function
	createProductRepoSamples(productRepo)
}

func createProductRepoSamples(productRepo *repository.MongoProductRepository) {
	err := services.CreateSampleProducts(productRepo)
	if err != nil {
		log.Fatal("Failed to create sample products:", err)
	}
}
