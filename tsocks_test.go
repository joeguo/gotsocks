package gotsocks

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	proxy, err := New("joe:abcd1234@198.98.108.119:54321", Socks5)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "wordpress.org")
	fmt.Println(string(result))
	if err != nil {
		t.Fatal(err)
	}

}

func aTestTimeout(t *testing.T){
	proxy, err := New("joe:abcd1234@199.188.74.170:54321", Socks5)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "yahoo.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(result))
}


