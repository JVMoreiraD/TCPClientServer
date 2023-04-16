package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"sdMatrix/matrix"
)

const (
	ConnHost = "localhost"
	ConnPort = "8080"
	ConnType = "tcp"
)

type Matrix [2][2]int

type Data struct {
	Matrix1 Matrix
	Matrix2 Matrix
}

func (d Data) MarshalJSON() ([]byte, error) {
	type Alias Data
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&d),
	})
}
func main() {
	l, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf, readErr := io.ReadAll(conn)
	var data Data

	if err := json.Unmarshal(buf, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	result := matrix.Multiply(matrix.Matrix(data.Matrix1), matrix.Matrix(data.Matrix2))

	if readErr != nil {
		fmt.Println("failed:", readErr)
		return
	}

	fmt.Println("Got: ", data, "\nResult: ", result)
	message, btl := json.Marshal(result)
	if btl != nil {
		fmt.Println("Error ao converter a structural para JSON:", btl)
		return
	}
	_, writeErr := conn.Write(message)
	if writeErr != nil {
		fmt.Println("failed:", writeErr)
		return
	}
	conn.Close()
}
