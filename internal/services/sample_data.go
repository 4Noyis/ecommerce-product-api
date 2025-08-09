package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/4Noyis/ecommerce-product-api/internal/models"
	"github.com/4Noyis/ecommerce-product-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSampleProducts returns a slice of sample product data for testing
func GetSampleProducts() []models.Product {
	now := time.Now()

	return []models.Product{
		{
			Name:        "iPhone 15 Pro",
			Description: "Latest Apple iPhone with advanced camera system and titanium design",
			Price:       999.99,
			Category:    "Electronics",
			Stock:       50,
			ImageURLs: []string{
				"https://example.com/iphone15pro-1.jpg",
				"https://example.com/iphone15pro-2.jpg",
			},
			Tags:   []string{"smartphone", "apple", "premium", "5g"},
			Brand:  "Apple",
			Weight: 0.187,
			Dimensions: models.Dimensions{
				Length: 14.67,
				Width:  7.09,
				Height: 0.83,
			},
			Attributes: map[string]string{
				"color":   "Natural Titanium",
				"storage": "128GB",
				"network": "5G",
			},
			Rating: models.ProductRating{
				Average: 4.5,
				Count:   120,
			},
			Status: "active",
		},
		{
			Name:        "MacBook Air M2",
			Description: "Thin, light laptop with M2 chip and all-day battery life",
			Price:       1199.99,
			Category:    "Computers",
			Stock:       25,
			ImageURLs: []string{
				"https://example.com/macbook-air-m2-1.jpg",
				"https://example.com/macbook-air-m2-2.jpg",
			},
			Tags:   []string{"laptop", "apple", "m2", "portable"},
			Brand:  "Apple",
			SKU:    "APPLE-MACBOOKAIRM2-001",
			Weight: 1.24,
			Dimensions: models.Dimensions{
				Length: 30.41,
				Width:  21.5,
				Height: 1.13,
			},
			Attributes: map[string]string{
				"color":   "Midnight",
				"storage": "256GB SSD",
				"memory":  "8GB",
				"chip":    "Apple M2",
			},
			Rating: models.ProductRating{
				Average: 4.7,
				Count:   85,
			},
			Status:    "active",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Name:        "Sony WH-1000XM5 Headphones",
			Description: "Premium noise-canceling wireless headphones with exceptional sound quality",
			Price:       399.99,
			Category:    "Audio",
			Stock:       75,
			ImageURLs: []string{
				"https://example.com/sony-wh1000xm5-1.jpg",
				"https://example.com/sony-wh1000xm5-2.jpg",
			},
			Tags:   []string{"headphones", "wireless", "noise-canceling", "premium"},
			Brand:  "Sony",
			SKU:    "SONY-WH1000XM5-001",
			Weight: 0.25,
			Dimensions: models.Dimensions{
				Length: 26.4,
				Width:  19.6,
				Height: 7.3,
			},
			Attributes: map[string]string{
				"color":        "Black",
				"connectivity": "Bluetooth 5.2",
				"battery":      "30 hours",
				"driver":       "30mm",
			},
			Rating: models.ProductRating{
				Average: 4.6,
				Count:   200,
			},
			Status:    "active",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Name:        "Nike Air Max 270",
			Description: "Comfortable lifestyle sneakers with Max Air cushioning",
			Price:       149.99,
			Category:    "Footwear",
			Stock:       100,
			ImageURLs: []string{
				"https://example.com/nike-airmax270-1.jpg",
				"https://example.com/nike-airmax270-2.jpg",
			},
			Tags:   []string{"sneakers", "casual", "comfort", "air-max"},
			Brand:  "Nike",
			SKU:    "NIKE-AIRMAX270-001",
			Weight: 0.8,
			Dimensions: models.Dimensions{
				Length: 31.0,
				Width:  11.5,
				Height: 12.0,
			},
			Attributes: map[string]string{
				"color":    "White/Black",
				"size":     "US 9",
				"material": "Mesh/Synthetic",
				"sole":     "Rubber",
			},
			Rating: models.ProductRating{
				Average: 4.3,
				Count:   150,
			},
			Status:    "active",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Name:        "Samsung 55\" 4K Smart TV",
			Description: "Crystal UHD 4K Smart TV with HDR and built-in streaming apps",
			Price:       699.99,
			Category:    "Electronics",
			Stock:       20,
			ImageURLs: []string{
				"https://example.com/samsung-55-4k-tv-1.jpg",
				"https://example.com/samsung-55-4k-tv-2.jpg",
			},
			Tags:   []string{"tv", "4k", "smart-tv", "hdr"},
			Brand:  "Samsung",
			SKU:    "SAMSUNG-55CU7000-001",
			Weight: 16.8,
			Dimensions: models.Dimensions{
				Length: 123.0,
				Width:  70.7,
				Height: 7.8,
			},
			Attributes: map[string]string{
				"screen_size": "55 inches",
				"resolution":  "3840 x 2160",
				"hdr":         "HDR10+",
				"os":          "Tizen",
			},
			Rating: models.ProductRating{
				Average: 4.2,
				Count:   95,
			},
			Status:    "active",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

// CreateSampleProducts inserts sample product data into the database
func CreateSampleProducts(productRepo *repository.MongoProductRepository) error {
	sampleProducts := GetSampleProducts()

	fmt.Println("Inserting sample products...")
	successCount := 0

	for i, product := range sampleProducts {
		err := productRepo.CreateProduct(&product)
		if err != nil {
			log.Printf("Failed to insert product %d (%s): %v", i+1, product.Name, err)
			continue
		}
		fmt.Printf("Inserted product: %s (SKU: %s)\n", product.Name, product.SKU)
		successCount++
	}

	if successCount == 0 {
		return fmt.Errorf("failed to insert any sample products")
	}

	fmt.Printf("\n Successfully inserted %d out of %d sample products!\n", successCount, len(sampleProducts))
	return nil
}

// GetSampleReviews returns a slice of sample review data for testing
func GetSampleReviews(productID primitive.ObjectID) []models.Review {
	now := time.Now()

	return []models.Review{
		{
			ProductID: productID,
			UserID:    "user123",
			Rating:    5,
			Title:     "Excellent product!",
			Comment:   "This product exceeded my expectations. The quality is outstanding and it arrived quickly. Highly recommended!",
			Verified:  true,
			Helpful:   15,
			CreatedAt: now.AddDate(0, 0, -7), // 7 days ago
			UpdatedAt: now.AddDate(0, 0, -7),
		},
		{
			ProductID: productID,
			UserID:    "user456",
			Rating:    4,
			Title:     "Great value for money",
			Comment:   "Really good product for the price. The only minor issue is the packaging could be better, but the product itself is solid.",
			Verified:  true,
			Helpful:   8,
			CreatedAt: now.AddDate(0, 0, -3), // 3 days ago
			UpdatedAt: now.AddDate(0, 0, -3),
		},
		{
			ProductID: productID,
			UserID:    "user789",
			Rating:    3,
			Title:     "Average product",
			Comment:   "It's okay, does what it's supposed to do. Nothing special but nothing wrong either. Would consider other options next time.",
			Verified:  false,
			Helpful:   2,
			CreatedAt: now.AddDate(0, 0, -1), // 1 day ago
			UpdatedAt: now.AddDate(0, 0, -1),
		},
		{
			ProductID: productID,
			UserID:    "user101",
			Rating:    5,
			Title:     "Perfect!",
			Comment:   "Absolutely love this! Works perfectly and the design is beautiful. Will definitely buy from this brand again.",
			Verified:  true,
			Helpful:   12,
			CreatedAt: now.AddDate(0, 0, -14), // 2 weeks ago
			UpdatedAt: now.AddDate(0, 0, -14),
		},
		{
			ProductID: productID,
			UserID:    "user202",
			Rating:    2,
			Title:     "Not as described",
			Comment:   "Product quality is below expectations. The description doesn't match what I received. Customer service was helpful though.",
			Verified:  true,
			Helpful:   3,
			CreatedAt: now.AddDate(0, 0, -5), // 5 days ago
			UpdatedAt: now.AddDate(0, 0, -5),
		},
	}
}

// CreateSampleReviews inserts sample review data into the database
func CreateSampleReviews(reviewRepo *repository.MongoReviewRepository, productID primitive.ObjectID) error {
	sampleReviews := GetSampleReviews(productID)
	ctx := context.Background()

	fmt.Printf("Inserting sample reviews for product ID: %s...\n", productID.Hex())
	successCount := 0

	for i, review := range sampleReviews {
		err := reviewRepo.Create(ctx, &review)
		if err != nil {
			log.Printf("Failed to insert review %d: %v", i+1, err)
			continue
		}
		fmt.Printf("Inserted review: %s (Rating: %d/5)\n", review.Title, review.Rating)
		successCount++
	}

	if successCount == 0 {
		return fmt.Errorf("failed to insert any sample reviews")
	}

	fmt.Printf("\n Successfully inserted %d out of %d sample reviews!\n", successCount, len(sampleReviews))
	return nil
}
