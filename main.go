package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("Listening on port :6379")

	//New server
	l,err := net.Listen("tcp" , ":6379")
	if err!=nil{
		fmt.Println(err)
		return
	}

	//Recieve requests
	conn,err := l.Accept()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//Receive commands and respond
    for{
	   resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
    }
}

