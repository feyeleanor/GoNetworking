package main
import "crypto/aes"
import "crypto/cipher"
import "fmt"
import "net"

const AES_KEY = "0123456789012345"

func main() {
	Dial(":1025", func(conn net.Conn) {
		RequestMessage(conn, func(m []byte) {
			if m, e := Decrypt(m, AES_KEY); e == nil {
				fmt.Printf("%s\n", m)
			}
		})
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

func RequestMessage(conn net.Conn, f func([]byte)) (e error) {
	if _, e = conn.Write([]byte("\n")); e == nil {
		m := make([]byte, 1024)
		var n int
		if n, e = conn.Read(m); e == nil {
			f(m[:n])
		}
	}
	return
}

func Decrypt(m []byte, k string) (r []byte, e error) {
	var b cipher.Block
	if b, e = aes.NewCipher([]byte(k)); e == nil {
		var iv []byte
		iv, m = Unpack(m)
		c := cipher.NewCBCDecrypter(b, iv)
		r = make([]byte, len(m))
		c.CryptBlocks(r, m)
	}
	return
}

func Unpack(m []byte) (iv, r []byte){
	return m[:aes.BlockSize], m[aes.BlockSize:]
}