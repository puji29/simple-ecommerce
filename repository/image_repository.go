package repository

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/entity"
	"log"
)

type ImageRepository interface {
	Create(payload entity.Images) (entity.Images, error)
	List() ([]entity.Images, error)
	GetImageById(id string) (entity.Images, error)
	UpdateImage(payload entity.Images) (entity.Images, error)
	DeleteImage(id string) (entity.Images, error)
}

type imageRepository struct {
	db *sql.DB
}

// DeleteImage implements ImageRepository.
func (i *imageRepository) DeleteImage(id string) (entity.Images, error) {
	_, err := i.db.Exec(config.DeleteImage, id)
	if err != nil {
		log.Println("repositoryQueryDelete:", err.Error())
		return entity.Images{}, err
	}
	return entity.Images{}, nil
}

// GetImageById implements ImageRepository.
func (i *imageRepository) GetImageById(id string) (entity.Images, error) {
	var image entity.Images

	err := i.db.QueryRow(config.SelectImageById, id).Scan(&image.ID, &image.Image, &image.CreatedAt, &image.UpdatedAt)
	if err != nil {
		log.Println("repositoryQuery.ById:", err.Error())
		return entity.Images{}, err
	}
	return entity.Images{}, nil
}

// UpdateImage implements ImageRepository.
func (i *imageRepository) UpdateImage(payload entity.Images) (entity.Images, error) {
	//bug
	err := i.db.QueryRow(config.ImageUpdate, payload.ID, payload.Image).Scan(&payload.UpdatedAt)
	if err != nil {
		log.Println("repositoryQuery:", err.Error())
		return entity.Images{}, err
	}
	return entity.Images{}, nil
}

// List implements ImageRepository.
func (i *imageRepository) List() ([]entity.Images, error) {
	var images []entity.Images

	rows, err := i.db.Query(config.SelectAllImage)
	if err != nil {
		log.Println("repositoryQueryList:", err.Error())
		return []entity.Images{}, err
	}

	for rows.Next() {
		var image entity.Images
		err := rows.Scan(
			&image.ID,
			&image.Image,
			&image.CreatedAt,
			&image.UpdatedAt,
		)
		if err != nil {
			log.Println("repositoryQuery.Next:", err.Error())
			return []entity.Images{}, err
		}
		images = append(images, image)
	}
	return images, nil
}

// Create implements ImageRepository.
func (i *imageRepository) Create(payload entity.Images) (entity.Images, error) {
	if err := i.db.QueryRow(config.InsertImage, payload.Image).Scan(&payload.ID, &payload.CreatedAt); err != nil {
		log.Println("Repository.Query Insert: ", err.Error())
		return entity.Images{}, err
	}
	return entity.Images{}, nil
}

func NewImageRepository(db *sql.DB) ImageRepository {
	return &imageRepository{db}
}
