package domain

import "time"

type EventStatus string

const (
	EventOpen    EventStatus = "OPEN"
	EventSoldOut EventStatus = "SOLD_OUT"
	EventClosed  EventStatus = "CLOSED"
)

type Venue struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Capacity int    `json:"capacity"`
}

type TicketType struct {
	ID            string  `json:"id"`
	EventID       string  `json:"event_id"`
	Name          string  `json:"name"`          // Ej: "VIP", "Campo"
	Price         float64 `json:"price"`
	StockAvailable int    `json:"stock_available"`
}

type Event struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Date        time.Time    `json:"date"`
	VenueID     string       `json:"venue_id"`
	Venue       *Venue       `json:"venue,omitempty"`
	TicketTypes []TicketType `json:"ticket_types,omitempty"`
	Status      EventStatus  `json:"status"`
	ImageURL    string       `json:"image_url"`
}