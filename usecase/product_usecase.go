package usecase

import (
	"ecommerce/entity"
	"ecommerce/repository"
	"log"
)

type ProductUseCase interface {
	CreateNewProduct(payload entity.Product) (entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	UpdateProduct(payload entity.Product) (entity.Product, error)
	DeletedProduct(id string) (entity.Product, error)
	FindProductById(id string) (entity.Product, error)
	FindProductByProductName(productName string) (entity.Product, error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

// FindProductByProductName implements ProductUseCase.
func (p *productUseCase) FindProductByProductName(productName string) (entity.Product, error) {
	return p.repo.GetProductByProductName(productName)
}

// FindProductById implements ProductUseCase.
func (p *productUseCase) FindProductById(id string) (entity.Product, error) {
	return p.repo.GetProductById(id)
}

// DeletedProduct implements ProductUseCase.
func (p *productUseCase) DeletedProduct(id string) (entity.Product, error) {
	return p.repo.DeleteProduct(id)
}

// UpdateProduct implements ProductUseCase.
func (p *productUseCase) UpdateProduct(payload entity.Product) (entity.Product, error) {
	if payload.ProductName == "" && payload.Description == "" && payload.Price < 0 && payload.StockQuantity < 0 && payload.CategoryId == "" && payload.ImageId == "" {
		log.Println("Field can't empty")
		return entity.Product{}, nil
	}
	return p.repo.UpdatedProduct(payload)
}

// FindAllProduct implements ProductUseCase.
func (p *productUseCase) FindAllProduct() ([]entity.Product, error) {
	return p.repo.ListProduct()
}

// CreateNewProduct implements ProductUseCase.
func (p *productUseCase) CreateNewProduct(payload entity.Product) (entity.Product, error) {
	return p.repo.Create(payload)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}
