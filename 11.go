package main
import "bufio"
import "fmt"
import "net"

func main() {
	if c, e := net.Dial("tcp", ":1024"); e == nil {
		defer c.Close()
		if m, e := bufio.NewReader(c).ReadString('\n'); e == nil {
			fmt.Printf(m)
		}
	}
}