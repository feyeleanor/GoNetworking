package main
import "crypto/rand"
import "crypto/tls"
import "fmt"
import "net"

func main() {
	Listen(":1025", ConfigTLS("scert", "skey"), func(c net.Conn) {
		fmt.Fprintln(c, "hello world")
	})
}

func ConfigTLS(c, k string) (r *tls.Config) {
	if cert, e := tls.LoadX509KeyPair(c, k); e == nil {
		r = &tls.Config{
			Certificates: []tls.Certificate{ cert },
			Rand: rand.Reader,
		}
	}
	return
}

func Listen(a string, conf *tls.Config, f func(net.Conn)) {
	if listener, e := tls.Listen("tcp", a, conf); e == nil {
		for {
			if connection, e := listener.Accept(); e == nil {
				go func(c net.Conn) {
					defer c.Close()
					f(c)
				}(connection)
			}
		}
	}
}