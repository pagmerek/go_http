package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func handleRequest(http_request map[string]string) {
	// switch http_request["request"] {
	// case "GET":
	// 	anwser, ok := words[resource]
	// 	if ok {
	// 		connection.Write([]byte("ANWSER " + anwser + "\n"))
	// 	} else {
	// 		connection.Write([]byte("ERROR can't find " + resource + "\n"))
	// 	}
	// case "PUT":
	// 	connection.Write([]byte("ADD DEFINITION FOR " + http_version + "\n"))
	// case "HEAD":

	// case "POST":

	// case "DELETE":
	// 	words = make(map[string]string)
	// case "OPTIONS":

	// default:
	// 	fmt.Println("INVALID HTTP METHOD")
	// }
}

func handleConnection(connection net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("New connection")
	reader := bufio.NewReader(connection)
	http_request := make(map[string]string)
	reply, _ := reader.ReadString('\n')
	http_request["request"] = reply
	for {
		reply, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			break
		}
		request := strings.Split(reply, ":")
		if len(request) > 1 {
			key, value := request[0], request[1]
			http_request[key] = value
		}
	}
	handleRequest(http_request)
}

func main() {
	portFlag := flag.String("p", "8000", "Please specify port for HTTP server")
	flag.Parse()
	fmt.Println("Starting server on http://127.0.0.1:" + *portFlag)
	socket, _ := net.Listen("tcp", ":"+*portFlag)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(l net.Listener) {
		for {
			c, err := l.Accept()
			if err != nil {
				return
			}
			wg.Add(1)
			go handleConnection(c, wg)
		}
	}(socket)
	time.Sleep(time.Minute)
	wg.Wait()
}
