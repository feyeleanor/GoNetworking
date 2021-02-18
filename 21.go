package main
import "bytes"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import"encoding/gob"
import "net"

func main() {
	HELLO_WORLD := []byte("Hello World")
	RSA_LABEL := []byte("served")
	Listen(":1025", func(c *net.UDPConn, a *net.UDPAddr, b []byte) {
		var key rsa.PublicKey
		if e := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&key); e == nil {
			if m, e := Encrypt(&key, HELLO_WORLD, RSA_LABEL); e == nil {
				c.WriteToUDP(m, a)
			}
		}
		return
	})
}

func Encrypt(key *rsa.PublicKey, m, l []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha1.New(), rand.Reader, key, m, l)
}

func Listen(a string, f func(*net.UDPConn, *net.UDPAddr, []byte)) {
	if address, e := net.ResolveUDPAddr("udp", a); e == nil {
		if conn, e := net.ListenUDP("udp", address); e == nil {
			for b := make([]byte, 1024); ; b = make([]byte, 1024) {
				if n, client, e := conn.ReadFromUDP(b); e == nil {
					go f(conn, client, b[:n])
				}
			}
		}
	}
	return
}