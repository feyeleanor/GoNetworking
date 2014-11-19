package main
import . "fmt"
import . "net/http"

func main() {
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, "hello world")
  })

  Spawn(
    func() { ListenAndServe(":1024", nil) },
    func() { ListenAndServeTLS(":1025", "cert.pem", "key.pem", nil) },
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