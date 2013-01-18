/*
Package gotsocks makes it easy to use tsocks proxy command easy in golang

A simple example:

package main

import (
"fmt"
"github.com/joeguo/gotsocks"
)

func main() {
    proxy, err := gotsocks.New("198.xx.xx.xx:12345", gotsocks.Socks5)
	if err != nil {
		fmt.Println(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
 */
package gotsocks

