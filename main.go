package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var words = make(map[string]string)

func main() {
	fmt.Println("Start server ...")
	socket, _ := net.Listen("tcp", ":8000")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(l net.Listener) {
		for {
			c, err := l.Accept()
			if err != nil {
				return
			}
			wg.Add(1)
			go func(connection net.Conn) {
				defer wg.Done()
				for {
					reader := bufio.NewReader(connection)
					for {
						reply, err := reader.ReadString('\n')
						if err != nil {
							fmt.Println("Connection closed")
							break
						}
						reply += " "
						request := strings.Split(reply, " ")
						switch request[0] {
						case "GET":
							anwser, ok := words[request[1]]
							if ok {
								connection.Write([]byte("ANWSER " + anwser + "\n"))
							} else {
								connection.Write([]byte("ERROR can't find " + request[1] + "\n"))
							}
						case "SET":
							words[request[1]] = strings.Join(request[2:], " ")
							connection.Write([]byte("ADD DEFINITION FOR " + request[1] + "\n"))
						case "CLEAR":
							words = make(map[string]string)
						case "ALL":
							for k, v := range words {
								connection.Write([]byte("WORD: " + k + " || DEFINITION: " + v + "\n"))
							}
						}
					}

				}
			}(c)
		}
	}(socket)
	time.Sleep(1 * time.Minute)
	wg.Wait()
}
