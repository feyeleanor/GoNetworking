package main
import "bytes"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "crypto/x509"
import "encoding/gob"
import "encoding/pem"
import "fmt"
import "io/ioutil"
import "net"

func main() {
  Dial(":1025", "ckey", func(c *net.UDPConn, k *rsa.PrivateKey) {
    if m, e := ReadStream(c); e == nil {
      if m, e := Decrypt(k, m, []byte("served")); e == nil {
        fmt.Println(string(m))
      }
    }
  })
}

func Dial(a, file string, f func(*net.UDPConn, *rsa.PrivateKey)) {
  if k, e := LoadPrivateKey(file); e == nil {
    if address, e := net.ResolveUDPAddr("udp", a); e == nil {
      if conn, e := net.DialUDP("udp", nil, address); e == nil {
        defer conn.Close()
        SendKey(conn, k.PublicKey, func() {
          f(conn, k)
        })
      }
    }
  }
}

func ReadStream(conn *net.UDPConn) (r []byte, e error) {
  m := make([]byte, 1024)
  var n int
  if n, e = conn.Read(m); e == nil {
    r = m[:n]
  }
  return
}

func Decrypt(key *rsa.PrivateKey, m, l []byte) ([]byte, error) {
  return rsa.DecryptOAEP(sha1.New(), rand.Reader, key, m, l)
}

func LoadPrivateKey(file string) (r *rsa.PrivateKey, e error) {
  if file, e := ioutil.ReadFile(file); e == nil {
    if block, _ := pem.Decode(file); block != nil {
      if block.Type == "RSA PRIVATE KEY" {
        r, e = x509.ParsePKCS1PrivateKey(block.Bytes)
      }
    }
  }
  return
}

func SendKey(c *net.UDPConn, k rsa.PublicKey, f func()) {
  var b bytes.Buffer
  if e := gob.NewEncoder(&b).Encode(k); e == nil {
    if _, e = c.Write(b.Bytes()); e == nil {
      f()
    }
  }
}