package model

// Data represents lines of data
type Data struct {
	Name     string `json:"name"`
	Quantity int    `json:"qty"`
	Price    int    `json:"price"`
}
