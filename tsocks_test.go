package gotsocks

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	proxy, err := New("198.xx.xx.xx:54345", Socks5)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "google.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(result))
}

func TestTimeout(t *testing.T){
	proxy, err := New("198.xx.xx.xx:54345", Socks5)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 1, "yahoo.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(result))
}


