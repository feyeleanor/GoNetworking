package main
import . "fmt"
import . "net/http"

func main() {
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, "hello world")
  })

  done := make(chan bool)
  go func() {
    ListenAndServe(":1024", nil)
    done <- true
  }()

  ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
  <- done
}