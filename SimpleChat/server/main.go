package main

import (
	"fmt"
	message "godemo/SimpleChat/model"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8765")
	defer listen.Close()
	if err != nil {
		fmt.Println("Fail to listen port 8765. Err:", err)
		return
	}
	for {
		fmt.Println("Server start to listen the port 8765")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Fail to accept. Error: ", err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {

	}
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 5000)
	_, readErr := conn.Read(buf[:4])
	if readErr != nil {
		fmt.Println("fail to get length of message. error is :", err)
		return
	}

	//var pkgLen int
	//binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	return
}
