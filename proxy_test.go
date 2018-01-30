package proxytools

import (
	"fmt"
	"testing"
)

func TestParseProxyFromUrl(t *testing.T) {
	proxy, e := ParseProxyFromUrl("http://gmail.com:8889")

	if e != nil {
		t.Errorf("ParseProxyFromUrl() error = %v", e)
		return
	}
	if proxy.host != "gmail.com" {
		t.Errorf("ParseProxyFromUrl() host = %v", proxy.host)
		return
	}
	if proxy.port != 8889 {
		t.Errorf("ParseProxyFromUrl() port = %v", proxy.port)
		return
	}
	if proxy.proto != "http" {
		t.Errorf("ParseProxyFromUrl() proto = %v", proxy.proto)
		return
	}
	fmt.Print(proxy)
}

func TestParseProxyFromUrlWithUnameLogin(t *testing.T) {
	proxy, e := ParseProxyFromUrl("http://test:test1@gmail.com:8889")
	if e != nil {
		t.Errorf("ParseProxyFromUrl() error = %v", e)
		return
	}
	if proxy.host != "gmail.com" {
		t.Errorf("ParseProxyFromUrl() host = %v", proxy.host)
		return
	}
	if proxy.port != 8889 {
		t.Errorf("ParseProxyFromUrl() port = %v", proxy.port)
		return
	}
	if proxy.proto != "http" {
		t.Errorf("ParseProxyFromUrl() proto = %v", proxy.proto)
		return
	}
	if proxy.login != "test" {
		t.Errorf("ParseProxyFromUrl() proto = %v", proxy.proto)
		return
	}
	if proxy.password != "test1" {
		t.Errorf("ParseProxyFromUrl() proto = %v", proxy.proto)
		return
	}

	fmt.Print(proxy)

	p, _ := ParseProxyFromUrl("http://test@gmail.com:8889")
	if proxy.login != "test" {
		t.Errorf("ParseProxyFromUrl() proto = %v", proxy.proto)
		return
	}
	fmt.Println(p)
}
