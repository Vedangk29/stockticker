package stock

import "time"

type StockUpdate struct {
	Symbol string
	Price  float64
	Time   time.Time
}
