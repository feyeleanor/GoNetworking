package main
import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "hello world")
	})

	done := make(chan bool)
	go func() {
		http.ListenAndServe(":1024", nil)
		done <- true
	}()

	http.ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
	<- done
}