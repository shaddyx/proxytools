package proxytools

import (
	"context"
	"net"
	"net/http"

	"golang.org/x/net/proxy"
)

type dialer struct {
	addr     string
	socks5   proxy.Dialer
	User     string
	Password string
}

func (d *dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	// TODO: golang.org/x/net/proxy need to add DialContext
	return d.Dial(network, addr)
}

func (d *dialer) Dial(network, addr string) (net.Conn, error) {
	var err error
	auth := &proxy.Auth{
		User:     d.User,
		Password: d.Password,
	}
	if len(d.User) == 0 {
		auth = nil
	}
	if d.socks5 == nil {
		d.socks5, err = proxy.SOCKS5("tcp", d.addr, auth, proxy.Direct)
		if err != nil {
			return nil, err
		}
	}
	return d.socks5.Dial(network, addr)
}

func Socks5Proxy(addr string, user string, password string) *http.Transport {
	d := &dialer{
		addr:     addr,
		User:     user,
		Password: password}
	return &http.Transport{
		DialContext: d.DialContext,
		Dial:        d.Dial,
	}
}

// func main() {
// 	http.DefaultTransport = Socks5Proxy("127.0.0.1:1080")

// 	resp, err := http.Get("http://www.google.com/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	io.Copy(os.Stdout, resp.Body)
// }
