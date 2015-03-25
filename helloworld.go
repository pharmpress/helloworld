package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"html/template"
	"hash/crc32"
	//"encoding/hex"
//	"os"
)

type Page struct {
    Title string
    Ip  string
    BgColor int
    FgColor int
}


var tmpl = `<html>
  <head>
    <title>{{.Title}}</title>
  </head>
  <body style="background-color:#{{printf "%x" .BgColor}};color:#{{printf "%x" .FgColor}}">
	<h1>{{.Title}}</h1>
    <p>
    	<ul>
    		<li>IP : {{.Ip}}</li>
 		<ul>
    </p>
  </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("foo").Parse(tmpl)
	ip, err := externalIP()
	
	hashCode := int(crc32.ChecksumIEEE([]byte(ip)))
	mask := 0x1000000
	bgcolor := hashCode%mask
	fgcolor := (mask - bgcolor -1)%mask

	p := &Page{Title: "Hello, World!", Ip: ip, BgColor: bgcolor, FgColor: fgcolor}
	err = t.Execute(w, p)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func main() {
	fmt.Println("Hello!")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("By!")
}