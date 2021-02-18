package main
import "fmt"
import "net/http"

const ADDRESS = ":1025"

func main() {
	message := "hello world"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
	})
	http.ListenAndServeTLS(ADDRESS, "cert.pem", "key.pem", nil)
}