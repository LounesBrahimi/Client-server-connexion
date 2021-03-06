package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:100")
	if err != nil {
		fmt.Println("Error of connexion")
		return
	}
	fmt.Println("Connected to localhost:100")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a message to the server")
	msg, _ := reader.ReadString('\n')
	fmt.Println("sending of : ", msg)
	fmt.Fprintf(conn, fmt.Sprint(msg))
}
