package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

//constants representing each type
const (
	STRING = '+'
    ERROR = '-'
	INTEGER = ':'
	BULK = '$'
	ARRAY = '*'
)

//struct to be used in the serialization and the deserialization process
type Value struct{
	typ string  //typ is used to determine the data type carried by the value.

	str string  //str holds the value of the string received from the simple strings.

	num int  // num holds the value of the integer received from the integers.

	bulk string //bulk is used to store the string received from the bulk strings.

	array []Value //array holds all the values received from the arrays.

}

type Resp struct{
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp{
    return &Resp{reader : bufio.NewReader(rd)}
}


// readLine() -> reads a single line from a RESP stream, stopping it after it encounters CRLF
func (r *Resp) readLine() (line []byte, n int, err error){
	for{
		b,err := r.reader.ReadByte()
		if err!=nil {
			return nil, 0,err
		}
		n+=1
		line = append(line,b)
		if len(line) >=2 && line[len(line)-2] == '\r' {
			break
		}

	}
	return line[:len(line)-2], n, nil
}

/*for example input : "$5\r\nAhmed\r\n", the readLine() function is called twice
  First call : 
    1) Reads : $,5,\r,\n
	2) Returns "$5", n=4
  Second call :
    1) Reads : A,h,m,e,d,\r,\n
	2) Returns : "Ahmed", n=7
*/


func (r*Resp) readInteger() (x int, n int, err error){
	line,n,err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line),10,64);

	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}


func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()

	if err != nil {
		return Value{}, err
	}

	switch _type {
	  case ARRAY:
		return r.readArray()
	  case BULK:
		return r.readBulk()
	  default:
		fmt.Printf("Unknown type: %v", string(_type))
		return Value{}, nil
	}
}

func (r *Resp) readArray() (Value, error){
	v := Value{}
	v.typ = "array"
	
	length,_,err := r.readInteger()

	if err != nil {
		return v, err
	}

	// foreach line, parse and read the value

	v.array = make([]Value, length)
	for i := 0; i<length; i++{
		val, err := r.Read()
		if err != nil{
			return v, err
		}
		// add parsed value to array
		v.array[i] = val
	}
	return v, nil
}


func (r *Resp) readBulk() (Value, error){
	v := Value{}
	v.typ = "bulk"

	len, _, err := r.readInteger()
	if err != nil{
		return v, err
	}

	bulk := make([]byte,len)

	r.reader.Read(bulk)

	v.bulk = string(bulk)
	//Read the trailing CRLF

	r.readLine()

	return v, nil
	
}