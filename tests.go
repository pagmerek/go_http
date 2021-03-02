package main

import (
	"bufio"
	"fmt"
	"net"
)

func tests() {
	test_coverage := 0
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("Can not connect to the server")
	} else {
		// todo: tests
		fmt.Println("TEST COVERAGE", test_coverage, "%")
	}
}
