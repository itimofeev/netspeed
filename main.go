package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	listen := flag.String("listen", "", "")
	connect := flag.String("connect", "", "")
	flag.Parse()

	if len(*listen) > 0 {
		doListen(*listen)
	} else if len(*connect) > 0 {
		doConnect(*connect)
	} else {
		panic("Neither listen nor connect not specified")
	}
}

func doConnect(addr string) {
	conn, err := net.Dial("tcp", addr)
	Must(err, "net.Dial")

	_, err = conn.Write([]byte("hello"))
	Must(err, "conn.Write")
	conn.Close()
}

func doListen(addr string) {
	l, err := net.Listen("tcp", addr)
	Must(err, "net.Listen")

	conn, err := l.Accept()
	Must(err, "l.Accept")

	data := make([]byte, 100)
	n, err := conn.Read(data)
	Must(err, "conn.Read")

	fmt.Println("!!!", string(data[:n]))
}

func Must(err error, msg ...string) {
	if err != nil {
		log.Panic(err, msg)
	}
}
