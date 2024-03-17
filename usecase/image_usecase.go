package usecase

import (
	"ecommerce/entity"
	"ecommerce/repository"
)

type ImageUseCase interface {
	CreateImage(payload entity.Images) (entity.Images, error)
	FindAllImage() ([]entity.Images, error)
	ImageUpdated(payload entity.Images) (entity.Images, error)
	ImageById(id string) (entity.Images, error)
	DeleteI(id string) (entity.Images, error)
}

type imageUseCase struct {
	repo repository.ImageRepository
}

// DeleteI implements ImageUseCase.
func (i *imageUseCase) DeleteI(id string) (entity.Images, error) {
	return i.repo.DeleteImage(id)
}

// ImageById implements ImageUseCase.
func (i *imageUseCase) ImageById(id string) (entity.Images, error) {
	return i.repo.GetImageById(id)
}

// ImageUpload implements ImageUseCase.
func (i *imageUseCase) ImageUpdated(payload entity.Images) (entity.Images, error) {
	return i.repo.UpdateImage(payload)
}

// FindAllImage implements ImageUseCase.
func (i *imageUseCase) FindAllImage() ([]entity.Images, error) {
	return i.repo.List()
}

// createImage implements ImageUseCase.
func (i *imageUseCase) CreateImage(payload entity.Images) (entity.Images, error) {
	return i.repo.Create(payload)
}

func NewImageUsecase(repo repository.ImageRepository) ImageUseCase {
	return &imageUseCase{repo}
}
