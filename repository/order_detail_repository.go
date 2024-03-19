package repository

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/entity"
	"log"
)

type OrderDetailRepository interface {
	Created(payload entity.OrderDetail) (entity.OrderDetail, error)
}

type orderDetailRepository struct {
	db *sql.DB
}

// created implements OrderDetailRepository.
func (o *orderDetailRepository) Created(payload entity.OrderDetail) (entity.OrderDetail, error) {
	if err := o.db.QueryRow(config.InsertOrderDetail, payload.OrderId, payload.ProductId, payload.Quantity, payload.TotalAmount).Scan(&payload.ID, &payload.CreatedAT); err != nil {
		log.Println("QueryRepositoryDetail: ", err.Error())
		return entity.OrderDetail{}, err
	}
	return entity.OrderDetail{}, nil
}

func NewOrderDetailRepository(db *sql.DB) OrderDetailRepository {
	return &orderDetailRepository{db}
}
