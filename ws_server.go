package main
import "fmt"
import "golang.org/x/net/websocket"
import "io/ioutil"
import "net/http"
import "strconv"

func SendJSON(ws *websocket.Conn, v ...interface{}) {
	websocket.JSON.Send(ws, v)
}

func ServeFile(route, name, mime_type string) {
	b, _ := ioutil.ReadFile(name)
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mime_type)
		fmt.Fprint(w, string(b))
	})
}

type Message struct {
	Author, Content string
}

type PigeonHole struct {
	*websocket.Conn
	Messages chan Message
	id int
}

func NewPigeonHole(ws *websocket.Conn, id int) *PigeonHole {
	return &PigeonHole { ws, make(chan Message, 16), id }
}

func (p *PigeonHole) Handshake() {
	websocket.JSON.Send(p.Conn, p.id)
}

func (p *PigeonHole) Listen(f func(r int, m Message)) {
	var b struct { Recipient, Content string }
	
	for {
		if e := websocket.JSON.Receive(p.Conn, &b); e == nil {
			r, _ := strconv.Atoi(b.Recipient)
			f(r, Message {
				Author: fmt.Sprint(p.id),
				Content: b.Content,
			})
		}
	}
}

func (p *PigeonHole) Deliver() {
fmt.Println("deliver outgoing messages")
	for {
		m := <- p.Messages
		SendJSON	(p.Conn, "private", m)
	}
}

func main() {
	h := NewPigeonHole(nil, 0)
	p := []*PigeonHole{ h }

	go func() {
		for {
			m := <- h.Messages
			go func() {
				for _, ph := range p[1:] {
					if ph.Conn != nil {
						SendJSON(ph.Conn, "broadcast", m)
					}
				}
			}()
		}
	}()

	http.Handle("/register", websocket.Handler(func(ws *websocket.Conn) {
		h := NewPigeonHole(ws, len(p))
		h.Handshake()
		p = append(p, h)

		go h.Deliver()
		h.Listen(func(r int, m Message) {
			if r < len(p) && r > -1 {
				p[r].Messages <- m
			}
		})
	}))

	ServeFile("/", "ws_server.html", "text/html")
	ServeFile("/js", "ws_server.js", "application/javascript")
	http.ListenAndServe(":3000", nil)
}
