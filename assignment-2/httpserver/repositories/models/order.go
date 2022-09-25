package models

import "time"

type Item struct {
	LineItemID uint `json:"line_time_id" gorm:"primary_key;autoIncrement"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity uint `json:"quantity"`
	OrderID uint `json:"-"`
}

type Order struct {
	OrderID       uint `json:"order_id" gorm:"primary_key;autoIncrement"`
	CustomerName  string `json:"customer_name"`
	OrderedAt time.Time `json:"ordered_at"`
	Items []Item `json:"items" gorm:"foreignKey:OrderID"`
}

