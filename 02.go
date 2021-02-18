package main
import "fmt"
import "net/http"

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
	http.HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, MESSAGE)
	})
	http.ListenAndServe(ADDRESS, nil)
}