package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "godemo/SimpleChat/model"
	"net"
)

func LoginIn(user ChatUser) (err error) {
	conn, err := net.Dial("tcp", "localhost:8765")
	if err != nil {
		fmt.Println("Fail to connect to server. Err: ", err, conn)
		return
	}
	defer conn.Close()

	var mess message.Message
	mess.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.Id = user.Id
	loginMes.Password = user.Password

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json format error. error = ", err)
	}
	mess.Data = string(data)

	realData, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("json format error. error = ", err)
	}

	pkgLen := uint32(len(realData))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	count, err := conn.Write(bytes[:4])
	if count != 4 || err != nil {
		fmt.Println("Fail to Write. Err= ", err)
	}

	fmt.Printf("length of message is %d \n content is %s", len(realData), realData)

	_, err = conn.Write(realData)
	if err != nil {
		fmt.Println("send message error. error = ", err)
	}

	return
}
