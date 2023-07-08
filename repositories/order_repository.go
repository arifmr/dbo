package repositories

import (
	"github.com/arifmr/dbo/database"
	"github.com/arifmr/dbo/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: database.DB,
	}
}

func (r *OrderRepository) GetOrders(offset, limit int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Offset(offset).Limit(limit).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	err := r.db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	err := r.db.Save(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) DeleteOrder(id int) error {
	err := r.db.Delete(&models.Order{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) SearchOrders(criteria string) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("product_name LIKE ?", "%"+criteria+"%").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
