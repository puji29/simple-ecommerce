package repository

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/entity"
	"log"
)

type ProductRepository interface {
	Create(payload entity.Product) (entity.Product, error)
	ListProduct() ([]entity.Product, error)
	UpdatedProduct(payload entity.Product) (entity.Product, error)
	DeleteProduct(id string) (entity.Product, error)
	GetProductById(id string) (entity.Product, error)
	GetProductByProductName(productName string) (entity.Product, error)
}

type productRepository struct {
	db *sql.DB
}

// GetProductByProductName implements ProductRepository.
func (p *productRepository) GetProductByProductName(productName string) (entity.Product, error) {
	var product entity.Product
	err := p.db.QueryRow(config.SelectProductByProductName, productName).Scan(&product.ID, &product.ProductName, &product.Description, &product.Price, &product.StockQuantity, &product.CreatedAT, &product.UpdatedAT, &product.CategoryId, &product.ImageId)
	if err != nil {
		log.Println("productRepository.GetByProductName:", err.Error())
		return entity.Product{}, err
	}
	return product, nil
}

// GetProductById implements ProductRepository.
func (p *productRepository) GetProductById(id string) (entity.Product, error) {
	var product entity.Product
	err := p.db.QueryRow(config.SelectProductById, id).Scan(&product.ID, &product.ProductName, &product.Description, &product.Price, &product.StockQuantity, &product.CreatedAT, &product.UpdatedAT, &product.CategoryId, &product.ImageId)
	if err != nil {
		log.Println("productRepository.GetById :", err.Error())
		return entity.Product{}, err
	}
	return product, nil
}

// DeleteProduct implements ProductRepository.
func (p *productRepository) DeleteProduct(id string) (entity.Product, error) {
	_, err := p.db.Exec(config.DeleteProduct, id)
	if err != nil {
		log.Println("query.Delete :", err.Error())
		return entity.Product{}, err
	}
	return entity.Product{}, nil
}

// UpdatedProduct implements ProductRepository.
func (p *productRepository) UpdatedProduct(payload entity.Product) (entity.Product, error) {
	if err := p.db.QueryRow(config.UpdateProduct, payload.ID, payload.ProductName, payload.Description, payload.Price, payload.StockQuantity, payload.CategoryId, payload.ImageId).Scan(&payload.UpdatedAT); err != nil {
		log.Println("queryProduct.Update :", err.Error())
		return entity.Product{}, err
	}
	return payload, nil
}

// ListProduct implements ProductRepository.
func (p *productRepository) ListProduct() ([]entity.Product, error) {
	var products []entity.Product

	rows, err := p.db.Query(config.SelectProduct)
	if err != nil {
		log.Println("productQuery:", err.Error())
		return []entity.Product{}, err
	}
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.StockQuantity,
			&product.CreatedAT,
			&product.UpdatedAT,
			&product.CategoryId,
			&product.ImageId,
		)
		if err != nil {
			log.Println("productRepository.Next: ", err.Error())
			return []entity.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

// create implements ProductRepository.
func (p *productRepository) Create(payload entity.Product) (entity.Product, error) {
	if err := p.db.QueryRow(config.InsertProduct, payload.ProductName, payload.Description, payload.Price, payload.StockQuantity, payload.CategoryId, payload.ImageId).Scan(&payload.ID, &payload.CreatedAT); err != nil {
		log.Println("queryInsertProduct :", err.Error())
		return entity.Product{}, err
	}
	return payload, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
