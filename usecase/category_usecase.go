package usecase

import (
	"ecommerce/entity"
	"ecommerce/repository"
)

type CategoryUseCase interface {
	InsertNew(payload entity.Category) (entity.Category, error)
}
type categoryUseCase struct {
	repo repository.CategoryRepository
}

// InsertNew implements CategoryUseCase.
func (c *categoryUseCase) InsertNew(payload entity.Category) (entity.Category, error) {
	return c.repo.Create(payload)
}

func NewCategoryUseCase(repo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{repo}
}
