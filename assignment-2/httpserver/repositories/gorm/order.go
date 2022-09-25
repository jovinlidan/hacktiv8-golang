package gorm

import (
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"

	"github.com/jinzhu/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (s *orderRepo) GetOrders() (*[]models.Order, error) {
	var orders []models.Order
	err := s.db.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (s *orderRepo) GetOrder(id string) (*models.Order, error) {
	var order models.Order
	err := s.db.Model(&models.Order{}).Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderRepo) CreateOrder(order *models.Order) error {
	err := s.db.Create(&order).Error
	return err
}

func (s *orderRepo) UpdateOrder(order *models.Order) (*models.Order, error) {
	err := s.db.Save(&order).Error
	return order ,err
}
func (s *orderRepo) DeleteOrder(order *models.Order) (*models.Order, error) {
	s.db.Where("order_id = ?", order.OrderID).Delete(&models.Item{})
	s.db.Where("order_id = ?", order.OrderID).Delete(&models.Order{})
	return order, nil
}

