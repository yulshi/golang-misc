package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("server failed to start: ", err)
		return
	}

	for {
		fmt.Println("Start to listen client's connections")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Warning: client connection can not be accepted: ", err)
		} else {
			go handle(conn)
		}
	}

}

func handle(conn net.Conn) {

	fmt.Printf("incoming request from %s\n", conn.RemoteAddr().String())
	defer func() {
		fmt.Println("connection is closing...")
		conn.Close()
		fmt.Println("connection is closed")
	}()

	buff := make([]byte, 512)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed the connection")
			} else {
				fmt.Println("Failed to read data: ", err)
			}
			return
		}

		text := string(buff[:n])
		//if strings.TrimSpace(text) == "exit" {
		//  return
		//}
		fmt.Printf("Received[%s]: %s", conn.RemoteAddr().String(), text)
	}

}
