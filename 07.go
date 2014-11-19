package main
import . "fmt"
import . "net/http"
import "sync"

var servers sync.WaitGroup

func main() {
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, "hello world")
  })

  Spawn(func() {
    ListenAndServe(":1024", nil)
  })
  Spawn(func() {
    ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil)
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