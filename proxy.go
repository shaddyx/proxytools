package proxytools

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	TYPE_SOCKS4 = "socks4"
	TYPE_SOCKS5 = "socks5"
	TYPE_HTTP   = "http"
)

type Proxy struct {
	proto    string
	url      string
	host     string
	port     int
	login    string
	password string
}

func defPortsMap(proto string) (int, error) {
	if proto == TYPE_SOCKS4 || proto == TYPE_SOCKS5 {
		return 1080, nil
	} else if proto == TYPE_HTTP {
		return 8080, nil
	} else {
		return 0, errors.New("Error, unknown protocol:" + proto)
	}
}
func parseLoginPass(urlChunk string) (string, string, error) {
	chunks := strings.SplitN(urlChunk, ":", 2)
	if len(chunks) == 1 {
		if len(chunks[0]) == 0 {
			return "", "", errors.New("Error parsing username and password")
		}
		return chunks[0], "", nil
	} else {
		return chunks[0], chunks[1], nil
	}
}

func splitHostPort(urlChunk string) (string, string, error) {
	chunks := strings.SplitN(urlChunk, ":", 2)
	if len(chunks) == 1 {
		if len(chunks[0]) == 0 {
			return "", "", errors.New("Error parsing hostName and port")
		}
		return chunks[0], "", nil
	} else {
		return chunks[0], chunks[1], nil
	}
}

func ParseProxyFromUrl(urlString string) (*Proxy, error) {
	// "socks5://127.0.0.1:9150"
	// tbProxyURL, err := url.Parse()

	chunks := strings.SplitN(urlString, "://", 2)
	p := new(Proxy)

	if len(chunks) == 1 {
		p.proto = TYPE_HTTP
	} else {
		p.proto = chunks[0]
		urlString = chunks[1]
	}

	chunks = strings.SplitN(urlString, "@", 2)
	if len(chunks) == 1 {
		urlString = chunks[0]
	} else {
		urlString = chunks[1]

		login, password, e := parseLoginPass(chunks[0])
		if e != nil {
			return nil, e
		}
		p.login = login
		p.password = password
	}
	host, port, err := splitHostPort(urlString)
	if err != nil {
		return nil, err
	}
	p.host = host
	if port != "" {
		iPort, err := strconv.Atoi(port)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("bad int:{} err:{}", port, err))
		}
		p.port = iPort
	} else {
		iPort, err := defPortsMap(p.proto)
		if err != nil {
			return nil, err
		}
		p.port = iPort
	}
	return p, nil
}

func HttpSetSockProxy(client *http.Client, proxy Proxy) (*http.Client, error) {
	client.Transport := Socks5Proxy {
		
	} 
}
func HttpSetProxy(client *http.Client, proxyUrl string) (*http.Client, error) {
	proxy, err := ParseProxyFromUrl(proxyUrl)
	if err != nil {
		return nil, err
	}
	if proxy.proto == TYPE_SOCKS5 {

	} else {

	}
	return client, nil
}
