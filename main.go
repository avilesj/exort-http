package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Print("Handle conn")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received: %s\n", message)
		conn.Close()
	}
}

func main() {
	fmt.Print("Launching server")
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Print("Server if online")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error on connection: %s", err)
		}

		go handleConn(conn)
	}
}
