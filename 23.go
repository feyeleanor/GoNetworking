package main
import "fmt"
import "golang.org/x/net/websocket"
import "io/ioutil"
import "net/http"

func main() {
	http.Handle("/hello", websocket.Handler(func(ws *websocket.Conn) {
		websocket.JSON.Send(ws, "Hello")
	}))
	ServeFile("/", "23.html", "text/html")
	ServeFile("/js", "23.js", "application/javascript")
	http.ListenAndServe(":3000", nil)
}

func ServeFile(route, name, mime_type string) {
	b, _ := ioutil.ReadFile(name)
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mime_type)
		fmt.Fprint(w, string(b))
	})
}