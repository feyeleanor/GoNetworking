package main
import "crypto/aes"
import "crypto/cipher"
import "net"

const AES_KEY = "0123456789012345"

func main() {
	Listen(":1025", func(c *net.UDPConn, a *net.UDPAddr, b []byte) {
		if m, e := Encrypt("Hello World", AES_KEY); e == nil {
			c.WriteToUDP(m, a)
		}
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

func Encrypt(m, k string) (o []byte, e error) {
	if o, e = PaddedBuffer([]byte(m)); e == nil {
		var b cipher.Block
		if b, e = aes.NewCipher([]byte(k)); e == nil {
			o = CryptBlocks(o, b)
		}
	}
	return
}

func PaddedBuffer(m []byte) (b []byte, e error) {
	b = append(b, m...)
	if p := len(b) % aes.BlockSize; p != 0 {
		p = aes.BlockSize - p
		b = append(b, make([]byte, p)...)  // padding with NUL!!!!
	}
	return
}

func CryptBlocks(b []byte, c cipher.Block) (o []byte) {
	o = make([]byte, aes.BlockSize + len(b))
	copy(o, IV())
	enc := cipher.NewCBCEncrypter(c, o[:aes.BlockSize])
	enc.CryptBlocks(o[aes.BlockSize:], b)
	return
}

func IV() (b []byte) {
	b = make([]byte, aes.BlockSize)
	if _, e := rand.Read(b); e != nil {
		panic(e)
	}
	return
}