package model

type ProductView struct {
	ID          int64       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"` 
	Category    string    `json:"category"`
	Photos		string		`json:"photos"`
	Comments    []Comment `json:"comments"`
}
