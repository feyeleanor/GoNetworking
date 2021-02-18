package main
import "fmt"
import "net/http"
import "sync"

var servers sync.WaitGroup

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "hello world")
	})

	Spawn(func() {
		http.ListenAndServe(":1024", nil)
	})
	Spawn(func() {
		http.ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
	})
	servers.Wait()
}

func Spawn(f ...func()) {
	for _, s := range f {
		servers.Add(1)
		go func(server func()) {
			defer servers.Done()
			server()
		}(s)
	}
}