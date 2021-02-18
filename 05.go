package main
import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "hello world")
	})

	Spawn(
		func() {
			http.ListenAndServe(":1024", nil)
		},
		func() {
			http.ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
		},
	)
}

func Spawn(f ...func()) {
	done := make(chan bool)
	for _, s := range f {
		go func(server func()) {
			server()
			done <- true
		}(s)
	}

	for l := len(f); l > 0; l-- {
		<- done
	}
}