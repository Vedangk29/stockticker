package display

import (
	"html/template"
	"net/http"
	"stockticker/stock"
	"sync"
)

var (
	mu      sync.Mutex
	updates = make(map[string]stock.StockUpdate)
	tmpl    = template.Must(template.ParseFiles("templates/index.html"))
)

func StartServer(ch <-chan stock.StockUpdate) {
	go func() {
		for update := range ch {
			mu.Lock()
			updates[update.Symbol] = update
			mu.Unlock()
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		tmpl.Execute(w, updates)
	})

	http.ListenAndServe(":8080", nil)
}
