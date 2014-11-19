package main
import . "fmt"
import "net"

func main() {
  Listen("tcp", ":1024", func(c net.Conn) {
    defer c.Close()
    Fprintln(c, "hello world")
  })
}

func Listen(p, a string, f func(net.Conn)) (e error) {
  var listener net.Listener
  if listener, e = net.Listen(p, a); e == nil {
    for {
      if connection, e := listener.Accept(); e == nil {
        go f(connection)
      }
    }
  }
  return
}