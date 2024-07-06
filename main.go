package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

const addr = "localhost:8080"
const network = "tcp4"

func main() {
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalf("Failed to listen network: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
		}

		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		str, err := reader.ReadBytes('\n')
		if err != nil {
			log.Printf("Failed to read input msg: %v", err)
			return
		}

		msg := strings.TrimSuffix(string(str), "\n")
		msg = strings.TrimSuffix(msg, "\r")

		res := "Получено сообщение: " + msg
		conn.Write([]byte(res + "\n"))
	}
}
