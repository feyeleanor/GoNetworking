package main
import "net"

func main() {
	HELLO_WORLD := []byte("Hello World\n")
	Listen(":1024", func(c *net.UDPConn, a *net.UDPAddr, b []byte) {
		c.WriteToUDP(HELLO_WORLD, a)
	})
}

func Listen(a string, f func(*net.UDPConn, *net.UDPAddr, []byte)) {
	if address, e := net.ResolveUDPAddr("udp", a); e == nil {
		if conn, e := net.ListenUDP("udp", address); e == nil {
			ServeUDP(conn, func(c *net.UDPAddr, b []byte) {
				f(conn, c, b)
			})
		}
	}
}

func ServeUDP(c *net.UDPConn, f func(*net.UDPAddr, []byte)) {
	for b := make([]byte, 1024); ; b = make([]byte, 1024) {
		if n, client, e := c.ReadFromUDP(b); e == nil {
			go f(client, b[:n])
		}
	}
}