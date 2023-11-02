package server
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)
func StartClient() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	// Read user's username
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	// Send the username to the server
	conn.Write([]byte(username))
	// Start reading user input and sending messages to the server
	go func() {
		for {
			msg, _ := reader.ReadString('\n')
			conn.Write([]byte(msg))
		}
	}()
	// Start reading messages from the server and printing them to the terminal
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(msg)
	}
}
