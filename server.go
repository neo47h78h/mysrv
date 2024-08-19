package main

import (
	"bufio"
	"fmt"
	"myserver/mods/ai"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	// Create a buffer to read the incoming data
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		// Read the client's message
		message := scanner.Text()
		fmt.Println("Received message from client:", message)

		// Respond to the client
		response, err := ai.GetAIResponse(message)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		_, _ = conn.Write([]byte(response + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from client:", err)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		os.Exit(1)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("fail to accept connection", err)
			continue
		}

		go handleConnection(conn)
	}
}
