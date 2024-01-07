package internals

import (
	MessageBox "TestRemote/Remote/Utility"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var exitFlag bool

func HandleConnection() {

	helpMessage := `
Available commands:
- h: Help Menu
- m: Call Package A API
- n: Call Package B API
- e: Exit the call
`

	for !exitFlag {
		// Read the input from the user
		fmt.Print("Enter a command (h for help): ")
		reader := bufio.NewReader(os.Stdin)
		clientInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim spaces and check the input
		clientInput = strings.TrimSpace(clientInput)

		switch clientInput {
		case "h":
			// Print the help message to the user
			fmt.Println(helpMessage)
		case "m":
			fmt.Println("You pressed 'm'")
			// Call the function to show the current directory
			// Example: MessageBox.PrintCurrDirectoryFromCmd()
		case "n":
			fmt.Println("You pressed 'n'")
			// Call the function for client connection
			MessageBox.ClientConn()
		case "e":
			fmt.Println("Exiting the app")
			exitFlag = true
		default:
			fmt.Println("Invalid input. No action taken.")
			break
		}
	}
}
