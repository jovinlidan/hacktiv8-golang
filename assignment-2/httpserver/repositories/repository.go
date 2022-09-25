package repositories

import "assignment2/httpserver/repositories/models"

type OrderRepo interface {
	CreateOrder(order *models.Order) error
	GetOrders() (*[]models.Order, error)
	UpdateOrder(order *models.Order) (*models.Order, error)
	DeleteOrder(order *models.Order) (*models.Order, error)
	GetOrder(id string)( *models.Order, error)
}
