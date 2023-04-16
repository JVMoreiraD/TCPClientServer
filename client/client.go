package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
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
	inputObject := Data{Matrix1: Matrix{}, Matrix2: Matrix{}}

	var rows, columns = 2, 2

	fmt.Print("Enter the First Matrix Items to Multiplication = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&inputObject.Matrix1[i][j])
		}
	}

	fmt.Print("Enter the Second Matrix Items to Multiplication = ")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&inputObject.Matrix2[i][j])
		}
	}

	conn, err := net.Dial(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	message, btl := json.Marshal(inputObject)
	if btl != nil {
		fmt.Println("Error ao converter a structural para JSON:", btl)
		return
	}

	ReadNWrite(conn, message)

	conn.Close()
}

func ReadNWrite(conn net.Conn, message []byte) {
	_, writeErr := conn.Write(message)
	if writeErr != nil {
		fmt.Println("failed:", writeErr)
		return
	}
	conn.(*net.TCPConn).CloseWrite()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Println("Result: ", string(buffer[:n]))
	}
}
