package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	PATCH  HttpMethod = "PATCH"
)

type Request struct {
	Version *string
	Path    *string
	Host    *string
	Agent   *string
	Method  HttpMethod
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Print("Handle conn")
	reader := bufio.NewReader(conn)
	var headers []string
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			fmt.Printf("Parsing error: %s", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		headers = append(headers, line)
	}

	println(headers)
	response := "HTTP/1.1 200 OK\r\n"
	conn.Write([]byte(response))
}

func main() {
	fmt.Println("Launching server")
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("Server if online")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error on connection: %s", err)
		}

		go handleConn(conn)
	}
}
