package gotsocks

import (
	"fmt"
	"bytes"
	"time"
	"strings"
	"errors"
	"os"
	"os/exec"
	"io/ioutil"
)

const (
	Socks5 = 5
	Socks4 = 4
)

var (
	ConfigLocation = `/tmp/tsocks-%s:%s.conf`
	PasswordConfig = `
local = 192.168.0.0/255.255.0.0
local = 127.0.0.1/255.0.0.0
local = %s.0.0.0/255.0.0.0
server = %s
server_type = %d
server_port = %s
default_user = %s
default_pass = %s`
	Config         = `
local = 192.168.0.0/255.255.0.0
local = 127.0.0.1/255.0.0.0
local = %s.0.0.0/255.0.0.0
server = %s
server_type = %d
server_port = %s`
)

type Proxy struct {
	IP         string
	Port       string
	User       string
	Password   string
	Version    int
}

//proxy format: user:password@ip:port
//version:Socks5 or Socks4
func New(proxy string, version int) (*Proxy, error) {
	temp := strings.FieldsFunc(proxy, func(r rune) bool {
			switch r{
			case '@', ':':
				return true
			}
			return false
		})
	switch len(temp){
	case 4:
		p := &Proxy{User:temp[0], Password:temp[1], IP:temp[2], Port:temp[3], Version:version}
		return p, p.sure()
	case 2:
		p := &Proxy{IP:temp[0], Port:temp[1], Version:version}
		return p, p.sure()

	}
	return nil, errors.New(fmt.Sprintf("Malformat proxy %s\n", proxy))
}
//execute command with tsocks
//timeout in seconds
func (proxy *Proxy) Command(name string, timeout int64, arg ...string) ([]byte, error) {
	var buffer bytes.Buffer
	var err error
	cmd := exec.Command(name, arg...)
	cmd.Env = []string{fmt.Sprintf("TSOCKS_CONF_FILE=%s", proxy.configFile())}
	cmd.Stdout = &buffer

	if err = cmd.Start(); err != nil {
		return buffer.Bytes(), err
	}
	c := make(chan error)
	go func() {
		c<- cmd.Wait()
	}()
	select {
	case res := <-c:
		err = res
	case <-time.After(time.Duration(timeout)*time.Second):
		cmd.Process.Kill()
		err = errors.New(fmt.Sprintf("%s timeout", name))
		return nil, err
	}
	return buffer.Bytes(), err

}

func (proxy *Proxy) configFile() string {
	return fmt.Sprintf(ConfigLocation, proxy.IP, proxy.Port)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return false
}

func (proxy *Proxy) sure() error {
	f := proxy.configFile()
	if !exists(f) {
		var content string
		ps := strings.Split(proxy.IP, ".")
		if proxy.User == "" {
			content = fmt.Sprintf(Config, ps[0], proxy.IP, proxy.Version, proxy.Port)
		}else {
			content = fmt.Sprintf(PasswordConfig, ps[0], proxy.IP, proxy.Version, proxy.Port, proxy.User, proxy.Password)
		}
		return ioutil.WriteFile(f, []byte(content), 0644)
	}
	return nil
}
