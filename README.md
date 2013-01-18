gotsocks
========

Execute command with [tsocks](http://tsocks.sourceforge.net)(a transparent SOCKS proxy library) in golang

Installation
==========
Install tsocks:
```sh
For Ubuntu/Debian

apt-get install tsocks

For CentOS/Fedora

yum install tsocks

```
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
		fmt.Println(err)
	}
	result, err := proxy.Command("/usr/bin/whois", 10, "google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
```
