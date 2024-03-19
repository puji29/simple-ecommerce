package repository

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/entity"
	"log"
	"strconv"
	"time"
)

type OrderTableRepository interface {
	Create(payload entity.OrderTable) (entity.OrderTable, error)
	List() ([]entity.OrderTable, error)
	Update(payload entity.OrderTable) (entity.OrderTable, error)
	Delete(id string) (entity.OrderTable, error)
	GetOrderTableByid(id string) (entity.OrderTable, error)
}

type orderTableRepository struct {
	db *sql.DB
}

// GetOrderTableByid implements OrderTableRepository.
func (o *orderTableRepository) GetOrderTableByid(id string) (entity.OrderTable, error) {
	var orderTable entity.OrderTable
	err := o.db.QueryRow(config.SelectOrderTableById, id).Scan(&orderTable.ID, &orderTable.UserId, &orderTable.OrderDate, &orderTable.TotalAmount, &orderTable.CreatedAt, &orderTable.UpdatedAt)
	if err != nil {
		log.Println("repositoryQueryUpdate:", err.Error())
		return entity.OrderTable{}, err
	}
	return orderTable, nil
}

// Delete implements OrderTableRepository.
func (o *orderTableRepository) Delete(id string) (entity.OrderTable, error) {
	_, err := o.db.Exec(config.DeleteOrderTable, id)
	if err != nil {
		log.Println("repositoryQuery.Delete:", err.Error())
		return entity.OrderTable{}, err
	}
	return entity.OrderTable{}, nil
}

// Update implements OrderTableRepository.
func (o *orderTableRepository) Update(payload entity.OrderTable) (entity.OrderTable, error) {
	payload.OrderDate = time.Now().Format("2006-01-02")
	if err := o.db.QueryRow(config.UpdateOrderItem, payload.ID, payload.OrderDate, payload.TotalAmount).Scan(&payload.UpdatedAt); err != nil {
		log.Println("repositoryUpdate:", err.Error())
		return entity.OrderTable{}, err
	}
	return entity.OrderTable{}, nil
}

// List implements OrderTableRepository.
func (o *orderTableRepository) List() ([]entity.OrderTable, error) {
	var orderTables []entity.OrderTable

	rows, err := o.db.Query(config.SelectOrderTable)
	if err != nil {
		log.Println("RepositoryOT.Query: ", err.Error())
		return []entity.OrderTable{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderTable entity.OrderTable
		var totalAmountStr string
		err := rows.Scan(
			&orderTable.ID,
			&orderTable.UserId,
			&orderTable.OrderDate,
			&totalAmountStr,
			&orderTable.CreatedAt,
			&orderTable.UpdatedAt,
		)
		if err != nil {
			log.Println("repositoryQuery.Next:", err.Error())
			return []entity.OrderTable{}, err
		}

		totalAmount, err := strconv.Atoi(totalAmountStr)
		if err != nil {
			log.Println("error conver")
			return []entity.OrderTable{}, err
		}
		orderTable.TotalAmount = totalAmount
		orderTables = append(orderTables, orderTable)
	}
	return orderTables, nil
}

// create implements OrderTableRepository.
func (o *orderTableRepository) Create(payload entity.OrderTable) (entity.OrderTable, error) {
	// today = time.Now()
	payload.OrderDate = time.Now().Format("2006-01-02")
	if err := o.db.QueryRow(config.InsertOrderTable, payload.UserId, payload.OrderDate, payload.TotalAmount).Scan(&payload.ID, &payload.CreatedAt); err != nil {
		log.Println("repostoryOT.QUERY : ", err.Error())
		return entity.OrderTable{}, err
	}
	return entity.OrderTable{}, nil
}

func NewOrderTableRepository(db *sql.DB) OrderTableRepository {
	return &orderTableRepository{db: db}
}
