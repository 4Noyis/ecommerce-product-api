package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
	Category    string             `json:"category" bson:"category"`
	Stock       int                `json:"stock" bson:"stock"`
	ImageURLs   []string           `json:"image_urls" bson:"image_urls"`
	Tags        []string           `json:"tags" bson:"tags"`
	Brand       string             `json:"brand" bson:"brand"`
	SKU         string             `json:"sku" bson:"sku"`
	Weight      float64            `json:"weight" bson:"weight"`
	Dimensions  Dimensions         `json:"dimensions" bson:"dimensions"`
	Attributes  map[string]string  `json:"attributes" bson:"attributes"`
	Rating      ProductRating      `json:"rating" bson:"rating"`
	Status      string             `json:"status" bson:"status"` // active, inactive
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type Dimensions struct {
	Length float64 `json:"length" bson:"length"`
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
}

type ProductRating struct {
	Average float64 `json:"average" bson:"average"`
	Count   int     `json:"count" bson:"count"`
}

type Review struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Rating    int                `json:"rating" bson:"rating"` // 1-5
	Title     string             `json:"title" bson:"title"`
	Comment   string             `json:"comment" bson:"comment"`
	Verified  bool               `json:"verified" bson:"verified"` // verified purchase
	Helpful   int                `json:"helpful" bson:"helpful"`   // helpful votes
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CartItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
	Price     float64            `json:"price" bson:"price"` // price at time of adding
	AddedAt   time.Time          `json:"added_at" bson:"added_at"`
}
