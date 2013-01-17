/*
Package gotsocks makes it easy to use tsocks proxy command easy in golang

A simple example:

package main

import (
"fmt"
"github.com/joeguo/gotsocks"
)

func main() {
    proxy, err := gotsocks.New("198.xx.xx.xx:12345", gotsocks.Socks5Version)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "google.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(result))
}
 */
package gotsocks

