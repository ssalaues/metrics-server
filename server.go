package main
import (
	"fmt"
	// "io"
	"log"
	"net/http"
	"os"
	"time"
)

const port string = "1234" // TODO get from env, default to 1234
const interval time.Duration = 1 // TODO make this configurable from env
type MetricsCache struct {
	Metrics	string
}

func ReadMetricsFromFile(file string) string {
	log.Println("Reading metrics file")
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err) // TODO possibly make soft fail and just log instead
	}
	return string(data)
}

func (m *MetricsCache) UpdateMetricsCache() chan string {
	ticker := time.NewTicker(interval * time.Second)
	quit := make(chan string)
	go func() {
		for {
			select {
				case <- ticker.C:
					m.Metrics = ReadMetricsFromFile("metrics_from_special_app.txt")
				case <- quit:
					ticker.Stop()
					return
			}
		}
	} ()
	return quit
}

func (m *MetricsCache) MetricsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("GET /metrics")
	fmt.Fprintf(w, m.Metrics)
}

func main() {
	log.Println("Metrics server starting up")

	// add new routes here
	metrics := &MetricsCache{""}
	metrics.UpdateMetricsCache()
	http.HandleFunc("/metrics", metrics.MetricsRoute)


	log.Println("Metrics server started on port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
