package sshat

import (
	"fmt"
	"os"
	"sshat/internal/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify 'server' or 'client' as a command-line argument.")
		return
	}

	switch os.Args[1] {
	case "--server":
		server.StartServer()
	case "--client":
		server.StartClient()
	default:
		fmt.Println("Invalid command-line argument. Please specify 'server' or 'client'.")
	}
}


