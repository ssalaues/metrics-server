package main
import (
	"fmt"
	// "io"
	"log"
	"net/http"
	"os"
)

const port string = "1234" // TODO get from env, default to 1234

func ReadMetricsFromFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err) // TODO possibly make soft fail and just log instead
	}
	return string(data)
}

func MetricsRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, ReadMetricsFromFile("metrics_from_special_app.txt"))
}

func main() {
	log.Println("Metrics server starting up")

	// add new routes here
	http.HandleFunc("/metrics", MetricsRoute)

	log.Println("Metrics server started on port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}