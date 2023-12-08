package main
import (
	"fmt"
	// "io"
	"log"
	"net/http"
	// "os"
)

const port string = "1234" // TODO get from env, default to 1234

func MetricsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Metric1\nMetric2\n")
}


func main() {
	log.Println("Metrics server starting up")

	http.HandleFunc("/metrics", MetricsRoute)

	log.Println("Metrics server started on port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}