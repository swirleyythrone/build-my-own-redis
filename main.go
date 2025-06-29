package main

import (
	"fmt"
	"io"
	"net"
	"os"
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
	   buf := make([]byte, 1024)// allocates a byte slice buffer to store the raw data read from the client over the TCP 
	   
	   //read message from the client
	   _,err = conn.Read(buf)
       
	   if err!=nil{
	   	if err == io.EOF {
               break // when the user inputs nothing break 
	   	}
	   	fmt.Println("Error reading from client: ",err.Error())
	   	os.Exit(1)
   
	   }
   
	   // ignore request and send back PONG
	   conn.Write([]byte("+OK\r\n"))
    }
}

