package usecase

import (
	"ecommerce/entity"
	"ecommerce/repository"
)

type OrderTableUseCase interface {
	CreateOrderTable(payload entity.OrderTable) (entity.OrderTable, error)
	FindAll() ([]entity.OrderTable, error)
	UpdateById(payload entity.OrderTable) (entity.OrderTable, error)
	DeleteOT(id string) (entity.OrderTable, error)
	FindByid(id string) (entity.OrderTable, error)
}

type orderTableUseCase struct {
	repo repository.OrderTableRepository
}

// FindByid implements OrderTableUseCase.
func (o *orderTableUseCase) FindByid(id string) (entity.OrderTable, error) {
	return o.repo.GetOrderTableByid(id)
}

// DeleteOT implements OrderTableUseCase.
func (o *orderTableUseCase) DeleteOT(id string) (entity.OrderTable, error) {
	return o.repo.Delete(id)
}

// UpdateById implements OrderTableUseCase.
func (o *orderTableUseCase) UpdateById(payload entity.OrderTable) (entity.OrderTable, error) {
	return o.repo.Update(payload)
}

// FindAll implements OrderTableUseCase.
func (o *orderTableUseCase) FindAll() ([]entity.OrderTable, error) {
	return o.repo.List()
}

// CreateOrderTable implements OrderTableUseCase.
func (o *orderTableUseCase) CreateOrderTable(payload entity.OrderTable) (entity.OrderTable, error) {
	return o.repo.Create(payload)
}

func NewOrderTableUseCase(repo repository.OrderTableRepository) OrderTableUseCase {
	return &orderTableUseCase{repo}
}
