package views

import "time"

type OrderGetAll struct {
	OrderID   uint `json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt time.Time `json:"ordered_at"`
	Items []OrderItemGetAll `json:"items"`
}

type OrderItemGetAll struct {
	LineItemID uint `json:"line_time_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity uint `json:"quantity"`
	OrderID uint `json:"-"`
}