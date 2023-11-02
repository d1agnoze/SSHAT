package server
import (
	"bufio"
	"fmt"
	"log"
	"net"
  "strings"
)

func StartServer() {
	// Create a TCP listener on a specific port (e.g., 8080)
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// Create a map to store connected clients
	clients := make(map[string]net.Conn)
	for {
		// Accept incoming connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// Read client's username
		username, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		// Store the client's connection
		clients[strings.TrimSpace(username)] = conn
		// Broadcast message to all connected clients
		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Println(err)
					break
				}
				fmt.Printf("%s: %s", username, msg)
				// Broadcast the message to all clients
				for clientName, clientConn := range clients {
					if clientName != strings.TrimSpace(username) {
						clientConn.Write([]byte(fmt.Sprintf("%s: %s", username, msg)))
					}
				}
			}
			// Remove client from the map when the connection is closed
			delete(clients, strings.TrimSpace(username))
			conn.Close()
		}()
	}
}
