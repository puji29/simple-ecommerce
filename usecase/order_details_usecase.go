package usecase

import (
	"ecommerce/entity"
	"ecommerce/repository"
)

type OrderDetailUsecase interface {
	InsertNew(payload entity.OrderDetail) (entity.OrderDetail, error)
}

type orderDetailUsecase struct {
	repo repository.OrderDetailRepository
}

// InsertNew implements OrderDetailUsecase.
func (o *orderDetailUsecase) InsertNew(payload entity.OrderDetail) (entity.OrderDetail, error) {
	return o.repo.Created(payload)
}

func NewOrderDeatailUsecase(repo repository.OrderDetailRepository) OrderDetailUsecase {
	return &orderDetailUsecase{repo: repo}
}
