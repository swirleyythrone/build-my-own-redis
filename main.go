package main

import {
	"fmt"
	"io"
	"net"
	"os"
}

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
	buf := make([]byte, 1024)// allocates a byte slice buffer to store the raw data read from the client over the TCP connection
}

