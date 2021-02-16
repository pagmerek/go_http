package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	words := make(map[string]string)

	fmt.Println("Start server ...")

	socket, _ := net.Listen("tcp", ":8000")
	connection, _ := socket.Accept()
	connection.Write([]byte("Connected ...\n"))
	for {
		reply, _ := bufio.NewReader(connection).ReadString('\n')
		request := strings.Split(reply, " ")
		switch request[0] {
		case "GET":
			anwser, ok := words[request[1]]
			if ok {
				connection.Write([]byte("ANWSER " + anwser))
			} else {
				connection.Write([]byte("ERROR can't find " + request[1]))
			}
		case "SET":
			words[request[1]] = strings.Join(request[1:], "")
		}
	}
}
