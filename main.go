package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"stockticker/display"
	"stockticker/fetcher"
	"stockticker/stock"
	"sync"
	"time"
)

func main() {
	tickers := []string{"RELIANCE", "TCS", "INFY", "HDFCBANK"}
	updateInterval := 5 * time.Second
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	updates := make(chan stock.StockUpdate)
	var wg sync.WaitGroup

	for _, symbol := range tickers {
		wg.Add(1)
		go fetcher.FetchStock(ctx, symbol, updateInterval, updates, &wg)
	}

	go display.StartServer(updates)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Shutting down...")
	cancel()
	wg.Wait()
}
