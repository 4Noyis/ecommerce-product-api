package services

import (
	"errors"
	"strings"

	"github.com/4Noyis/ecommerce-product-api/internal/models"
	"github.com/4Noyis/ecommerce-product-api/internal/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepo repository.MongoProductRepository
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	// Validation
	if err := s.validateProduct(product); err != nil {
		return err
	}

	// Generate SKU if not provided
	if product.SKU == "" {
		product.SKU = s.generateSKU(product.Name, product.Brand)
	}

	// Initialize rating
	product.Rating = models.ProductRating{Average: 0, Count: 0}

	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) validateProduct(product *models.Product) error {
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if strings.TrimSpace(product.Category) == "" {
		return errors.New("product category is required")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	return nil
}

func (s *ProductService) generateSKU(name, brand string) string {
	// Simple SKU generation: BRAND-NAME-UUID
	cleanName := strings.ReplaceAll(strings.ToUpper(name), " ", "")
	cleanBrand := strings.ReplaceAll(strings.ToUpper(brand), " ", "")
	uniqueID := uuid.New().String()[:8]

	if cleanBrand == "" {
		return cleanName + "-" + uniqueID
	}

	return cleanBrand + "-" + cleanName + "-" + uniqueID
}
