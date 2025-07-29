package fetcher

import (
	"context"
	"math/rand"
	"stockticker/stock"
	"sync"
	"time"
)

func getMockINRPrice(symbol string) float64 {
	basePrices := map[string]float64{
		"RELIANCE": 2800.50,
		"TCS":      3600.25,
		"INFY":     1550.10,
		"HDFCBANK": 1650.75,
	}
	base := basePrices[symbol]
	variance := base * 0.01
	return base + (rand.Float64()-0.5)*2*variance
}

func FetchStock(ctx context.Context, symbol string, interval time.Duration, ch chan<- stock.StockUpdate, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			price := getMockINRPrice(symbol)
			ch <- stock.StockUpdate{
				Symbol: symbol,
				Price:  price,
				Time:   time.Now(),
			}
		}
	}
}
