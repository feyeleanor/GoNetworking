package main
import "bufio"
import "crypto/tls"
import . "fmt"
import "net"
 
func main() {
  Dial(":1025", ConfigTLS("ccert", "ckey"), func(c net.Conn) {
    if m, e := bufio.NewReader(c).ReadString('\n'); e == nil {
      Printf(m)
    }
  })
}

func ConfigTLS(c, k string) (r *tls.Config) {
  if cert, e := tls.LoadX509KeyPair(c, k); e == nil {
    r = &tls.Config{
      Certificates: []tls.Certificate{ cert },
      InsecureSkipVerify: true,
    }
  }
  return
}

func Dial(a string, conf *tls.Config, f func(net.Conn)) {
  if c, e := tls.Dial("tcp", a, conf); e == nil {
    defer c.Close()
    f(c)
  }
}