package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"sync"
	"time"
)

func handleRequest(http_request map[string]string, connection net.Conn) {
	request := strings.Split(http_request["request"], " ")
	fmt.Println(request)
	method, resource, http_version := request[0], request[1], request[2]
	switch method {
	case "GET":

		file, err := ioutil.ReadFile("www/" + resource)
		if err != nil {
			//status_code = 404
			connection.Write([]byte(http_version + " 404 NOT FOUND"))
			fmt.Println("404")
			return
		}
		connection.Write([]byte("HTTP/1.1 200 OK"))
		connection.Write([]byte("\r\n"))
		fmt.Println(string(file))
		connection.Write(file)

	case "PUT":
	case "HEAD":
	case "POST":
	case "DELETE":
	case "OPTIONS":
	default:
		fmt.Println("INVALID HTTP METHOD")
	}
}

func handleConnection(connection net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("New connection")
	reader := bufio.NewReader(connection)
	http_request := make(map[string]string)
	reply, err := reader.ReadString('\n')
	http_request["request"] = reply
	for reply != "\r\n" {
		reply, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		request := strings.Split(reply, ":")
		if len(request) > 1 {
			key, value := request[0], request[1]
			http_request[key] = value
		}
	}
	handleRequest(http_request, connection)
	connection.Close()
	fmt.Println("Connection closed")
}
func startServer(port string) {
	fmt.Println("Starting server on http://127.0.0.1:" + port)
	connection, _ := net.Listen("tcp", ":"+port)
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
	}(connection)
	time.Sleep(time.Minute)
	wg.Wait()
}
func main() {
	portFlag := flag.String("p", "8000", "Please specify port for HTTP server")
	flag.Parse()
	startServer(*portFlag)
}
