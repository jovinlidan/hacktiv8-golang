package params

import "time"

type OrderCreateRequest struct {
	CustomerName string `validate:"required" json:"customer_name"`
	OrderedAt time.Time `validate:"required" json:"ordered_at"`
	Items []OrderItemCreateRequest `validate:"required" json:"items" `
}



type OrderItemCreateRequest struct{
	ItemCode string `validate:"required" json:"item_code"`
	Description string `validate:"required" json:"description"`
	Quantity uint `validate:"required" json:"quantity"`
}

type OrderUpdateRequest struct {
	OrderID uint `validate:"required" json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt time.Time `json:"ordered_at"`
	Items []OrderItemUpdateRequest `json:"items" `
}
type OrderItemUpdateRequest struct{
	LineItemID uint `json:"line_time_id"`
	ItemCode string `json:"item_code"`
	Description string `json:"description"`
	Quantity uint `json:"quantity"`

}
