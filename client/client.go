// TCP Client in Go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	for {
		fmt.Print("Enter message: ")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		reply, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading reply:", err)
			return
		}
		fmt.Printf("Server reply: %s", reply)
	}
}
