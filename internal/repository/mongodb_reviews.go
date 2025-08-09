package repository

import (
	"context"

	"github.com/4Noyis/ecommerce-product-api/internal/models"
	"github.com/4Noyis/ecommerce-product-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoReviewRepository struct {
	db *database.MongoDB
}

func NewMongoReviewsRepository(db *database.MongoDB) *MongoReviewRepository {
	return &MongoReviewRepository{db: db}
}

func (r *MongoReviewRepository) Create(ctx context.Context, review *models.Review) error {
	_, err := r.db.ReviewsCollection().InsertOne(ctx, review)
	return err
}

func (r *MongoReviewRepository) GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]*models.Review, error) {
	cursor, err := r.db.ReviewsCollection().Find(ctx, map[string]interface{}{"product_id": productID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []*models.Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}
