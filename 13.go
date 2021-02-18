package main
import "crypto/rand"
import "crypto/tls"
import "fmt"

func main() {
	Listen(":1025", ConfigTLS("scert", "skey"), func(c *tls.Conn) {
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

func Listen(a string, conf *tls.Config, f func(*tls.Conn)) {
	if listener, e := tls.Listen("tcp", a, conf); e == nil {
		for {
			if connection, e := listener.Accept(); e == nil {
				go func(c *tls.Conn) {
					defer c.Close()
					f(c)
				}(connection.(*tls.Conn))
			}
		}
	}
}