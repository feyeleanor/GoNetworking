package main
import "fmt"
import "golang.org/x/net/websocket"

const SERVER = "ws://localhost:3000/hello"
const ADDRESS = "http://localhost/"

func main() {
	if ws, e := websocket.Dial(SERVER, "", ADDRESS); e == nil {
		var s string
		if e := websocket.JSON.Receive(ws, &s); e == nil {
			fmt.Println(s)
		}
	}
}