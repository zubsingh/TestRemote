package main

import (
	"TestRemote/Remote/internals"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const (
	port = ":4444"
)

var wg sync.WaitGroup

func main() {
	//MessageBox.MessageBox1()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	// Handle termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Wait for a termination signal
		<-stop

		// Close the listener and exit
		fmt.Println("\nReceived termination signal. Closing the server.")
		listener.Close()
		os.Exit(0)
	}()

	fmt.Println("Server is listening on", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			// Check if the listener was closed intentionally
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}

			fmt.Println("Error accepting connection:", err)
			continue
		}
		wg.Add(1)
		go internals.HandleConnection(conn, &wg)
	}
	wg.Wait()
}
