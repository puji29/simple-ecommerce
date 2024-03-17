package repository

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/entity"
	"log"
)

type CategoryRepository interface {
	Create(payload entity.Category) (entity.Category, error)
}

type categoryRepository struct {
	db *sql.DB
}

// Create implements CategoryRepository.
func (c *categoryRepository) Create(payload entity.Category) (entity.Category, error) {
	if err := c.db.QueryRow(config.InsertCategory, payload.CategoryName).Scan(&payload.ID, &payload.CreatedAt); err != nil {
		log.Println("ReositoryQueryInsert: ", err.Error())
		return entity.Category{}, err
	}
	return entity.Category{}, nil
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db}
}
