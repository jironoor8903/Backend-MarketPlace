package model

type Product struct {
	ID          int64       `json:"id"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
}
