package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	test_coverage := 0
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("Can not connect to the server")
	} else {
		conn.Write([]byte("SET foo temp\n"))
		reader := bufio.NewReader(conn)
		reply, _ := reader.ReadString('\n')
		if reply != "ADD DEFINITION FOR foo\n" {
			fmt.Println("SET ERROR")
		} else {
			test_coverage += 25
		}
		fmt.Println("TEST COVERAGE:", test_coverage, "%")
		conn.Write([]byte("GET foo \n"))
		reply, _ = reader.ReadString('\n')
		if reply != "ANWSER temp\n" {
			fmt.Println("GET ERROR")
		} else {
			test_coverage += 25
		}
		fmt.Println("TEST COVERAGE:", test_coverage, "%")

		conn.Write([]byte("ALL \n"))
		reply, _ = reader.ReadString('\n')
		reply, _ = reader.ReadString('\n')

		if reply != "WORD: foo || DEFINITION: temp\n" {
			fmt.Println("ALL ERROR")
		} else {
			test_coverage += 25
		}
		fmt.Println("TEST COVERAGE:", test_coverage, "%")
		conn.Write([]byte("CLEAR \n"))
		conn.Write([]byte("GET foo \n"))
		reply, _ = reader.ReadString('\n')
		reply, _ = reader.ReadString('\n')
		if reply != "ERROR can't find foo\n" {
			fmt.Println("CLEAR ERROR")
		} else {
			test_coverage += 25
		}
		fmt.Println("TEST COVERAGE", test_coverage, "%")
	}
}
