package services

import (
	"firstproject/models"
	"firstproject/repositories"
)

type ProductService interface {
	AddProduct(product *models.Product) error
	GetProducts() ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID string) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) AddProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) GetProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(productID string) error {
	return s.repo.DeleteProduct(productID)
}
