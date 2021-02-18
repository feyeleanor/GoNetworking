package main
import "fmt"
import "golang.org/x/net/websocket"
import "log"

const SERVER = "ws://localhost:3000/register"
const ADDRESS = "http://localhost/"

func main() {
	if ws, e := websocket.Dial(SERVER, "", ADDRESS); e == nil {
		var id int
		if e := websocket.JSON.Receive(ws, &id); e == nil {
			fmt.Println("Connected as user", id)

			var pub_count, priv_count int
			for {
				var m []interface{}
				switch e := websocket.JSON.Receive(ws, &m); {
				case e != nil:
					log.Fatal(e)

				case m[0] == "broadcast":
					pub_count++
					printMessage("BROADCAST", pub_count, m[1])

				case m[0] == "private":
					priv_count++
					printMessage("PRIVATE", priv_count, m[1])
				}
			}
		}
	}
}

func printMessage(t string, n int, m interface{}) {
	if x, ok := m.(map[string] interface{}); ok {
		fmt.Printf("*** %v (%v) from %v ***\n", t, n, x["Author"])
		fmt.Printf("\t%v\n", x["Content"])
	} else {
		fmt.Printf("*** UNKNOWN ***\n\t%v\n", m)
	}
}