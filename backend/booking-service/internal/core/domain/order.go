package domain

import "time"

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING" 
	OrderConfirmed OrderStatus = "CONFIRMED" 
	OrderCancelled OrderStatus = "CANCELLED" 
)

type OrderItem struct {
	ID           string  `json:"id"`
	OrderID      string  `json:"order_id"`
	EventID      string  `json:"event_id"`
	TicketTypeID string  `json:"ticket_type_id"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
}

type Order struct {
	ID          string      `json:"id"`
	UserID      string      `json:"user_id"`    // Solo el ID del usuario de Auth
	TotalAmount float64     `json:"total_amount"`
	Status      OrderStatus `json:"status"`
	Items       []OrderItem `json:"items"`
	CreatedAt   time.Time   `json:"created_at"`
}