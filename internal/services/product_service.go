package services

import (
	"learning-golang/internal/models"
	"learning-golang/internal/repository"
)

type ProductService interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProducts() ([]models.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepo}
}

func (s *productService) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *productService) GetProducts() ([]models.Product, error) {
	return s.productRepository.GetProducts()
}
