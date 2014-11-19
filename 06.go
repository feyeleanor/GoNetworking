package main
import . "fmt"
import . "net/http"
import "sync"

func main() {
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {})

  var servers sync.WaitGroup
  servers.Add(1)
  go func() {
    defer servers.Done()
    ListenAndServe(":1024", nil)
  }()

  servers.Add(1)
  go func() {
    defer servers.Done()
    ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
  }()
  servers.Wait()
}