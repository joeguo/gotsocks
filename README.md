gotsocks
========

Execute command with [tsocks](http://tsocks.sourceforge.net) proxy in golang

Installation
==========
Install gotsocks:

```sh
go get github.com/joeguo/gotsocks

```
Example
==========
```go
package main

import (
"fmt"
"github.com/joeguo/gotsocks"
)

func main() {
    proxy, err := gotsocks.New("198.xx.xx.xx:12345", gotsocks.Socks5)
	if err != nil {
		t.Fatal(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "google.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(result))
}
```
