package main
import "bufio"
import "fmt"
import "net"

func main() {
	Dial("tcp", ":1024", func(c net.Conn) {
		if m, e := bufio.NewReader(c).ReadString('\n'); e == nil {
			fmt.Printf(m)
		}
	})
}

func Dial(p, a string, f func(net.Conn)) (e error) {
	var c net.Conn
	if c, e = net.Dial("tcp", ":1024"); e == nil {
		defer c.Close()
		f(c)
	}
	return
}