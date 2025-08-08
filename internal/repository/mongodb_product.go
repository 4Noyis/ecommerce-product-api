package repository

import (
	"context"
	"errors"
	"time"

	"github.com/4Noyis/ecommerce-product-api/internal/models"
	"github.com/4Noyis/ecommerce-product-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoProductRepository struct {
	db *database.MongoDB
}

func NewMongoProductRepository(db *database.MongoDB) *MongoProductRepository {
	return &MongoProductRepository{db: db}
}

func (r *MongoProductRepository) CreateProduct(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if product.Status == "" {
		product.Status = "active"
	}

	_, err := r.db.ProductsCollection().InsertOne(ctx, product)
	return err
}

func (r *MongoProductRepository) GetProduct(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product *models.Product
	err = r.db.ProductsCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return product, nil

}

func (r *MongoProductRepository) UpdateProduct(product *models.Product) error {
	return errors.New("error")
}
