package domain

import (
	"time"
)

type Event struct {
	ID          string 		`json:"id"`
	Title       string    	`json:"title"`
	Description *string   	`json:"description"`
	Image       string    `json:"image"`
    Category    string    `json:"category"`
    Price       float64   `json:"price"`
	Location    string    	`json:"location"`
	EventDate   time.Time 	`json:"event_date"`
	Capacity    int32     	`json:"capacity"`
}

type GetEventsResponse struct {
	Events	[]Event	`json:"events"`
}