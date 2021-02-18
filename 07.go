package main
import "net/http"
import "sync"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {})

	var servers sync.WaitGroup
	servers.Add(1)
	go func() {
		defer servers.Done()
		http.ListenAndServe(":1024", nil)
	}()

	servers.Add(1)
	go func() {
		defer servers.Done()
		http.ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
	}()
	servers.Wait()
}