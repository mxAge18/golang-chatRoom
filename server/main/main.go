package main

import (
	"fmt"
	"net"
)
func process(conn net.Conn) {
	defer conn.Close()
	subProcessor := Processor{
		Conn : conn,
	}
	subProcessor.SubProcessor()

}
func main() {
	fmt.Println("start listen on the port 8888")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen error is", err)
		return
	}
	for {
		fmt.Println("start accpet the connection")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() error is", err)
			// return
		}
		go process(conn)
	}

}
