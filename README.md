**Real-Time Stock Market Ticker (INR) — Built with Go:**

A real-time stock market ticker built in Golang that simulates INR-based stock price updates for Indiancompanies using Goroutines and Channels. It features a colorful, responsive web UI with search and auto-refresh, ideal for learning concurrency and live data rendering in Go.

**Key Concepts:**

- Goroutines for parallel stock fetching
- Channel-based data communication
- Context handling for graceful shutdown
- Modular Go project architecture
- Live-updating HTML UI with Go templating

---

**Features :**

- Simulated INR stock prices
- Concurrent fetching using Goroutines
- Graceful shutdown via `context.Context`
- Dark-themed responsive web UI
- Auto-refresh every 5 seconds
- Search/filter by stock symbol

---

**Project Structure**

stockticker/
├── main.go # Entry point & signal handling
├── fetcher/ # Goroutines simulating price fetches
│ └── fetcher.go
├── stock/ # Shared StockUpdate struct
│ └── stock.go
├── display/ # Web server and HTML rendering
│ └── display.go
└── templates/
└── index.html # Responsive UI for stock data

yaml
Copy
Edit

---

**Sample Companies (Simulated)**

| Symbol     | Example Company Name        |
|------------|-----------------------------|
| RELIANCE   | Reliance Industries Ltd     |
| TCS        | Tata Consultancy Services   |
| INFY       | Infosys Ltd                 |
| HDFCBANK   | HDFC Bank Ltd               |

_You can easily extend this by adding more entries to the fetcher module._

---

**How to Run**

1. Clone the repo:
   ```bash
   git clone https://github.com/your-username/stockticker.git
   cd stockticker
Initialize the Go module:

bash
Copy
Edit
go mod init stockticker
go mod tidy
Run the project:

bash
Copy
Edit
go run main.go
Open your browser:

arduino
Copy
Edit
http://localhost:8080

**Live UI Preview:**

The UI refreshes every 5 seconds and highlights stock prices and last update timestamps with stylish formatting.

Dark mode

Search box for stock filtering

Mobile responsive layout


**Learning Outcomes:**

Real-world use of Goroutines and channels

Building responsive web apps in Go

Understanding concurrency and clean shutdown

Working with templates and dynamic HTML rendering

