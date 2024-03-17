package usecase

import (
	"ecommerce/entity"
	"ecommerce/repository"
)

type UserUseCase interface {
	RegisterNewUser(payload entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	Update(data entity.User) (entity.User, error)
	Delete(id string) (entity.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// Delete implements UserUseCase.
func (u *userUseCase) Delete(id string) (entity.User, error) {
	return u.repo.Delete(id)
}

// Update implements UserUseCase.
func (u *userUseCase) Update(data entity.User) (entity.User, error) {
	return u.repo.UpdateUser(data)
}

// FindAll implements UserUseCase.
func (u *userUseCase) FindAll() ([]entity.User, error) {
	return u.repo.List()
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload entity.User) (entity.User, error) {
	return u.repo.Create(payload)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
