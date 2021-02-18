package main
import "bufio"
import "fmt"
import "net"

func main() {
	Dial(":1024", func(conn net.Conn) {
		if _, e := conn.Write([]byte("\n")); e == nil {
			if m, e := bufio.NewReader(conn).ReadString('\n'); e == nil {
				fmt.Printf("%v", m)
			}
		}
	})
}

func Dial(a string, f func(net.Conn)) {
	if address, e := net.ResolveUDPAddr("udp", a); e == nil {
		if conn, e := net.DialUDP("udp", nil, address); e == nil {
			defer conn.Close()
			f(conn)
		}
	}
}