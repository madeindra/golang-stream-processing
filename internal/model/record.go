package model

// Data represents lines of data
type Data struct {
	Name     string `json:"name"`
	Quantity int    `json:"qty"`
	Price    int    `json:"price"`
}

// StockDetail represents detail of stock
type StockDetail struct {
	Name         string
	Quantity     int
	AveragePrice float64
}
