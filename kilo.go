package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func enableRawMode() *term.State {
	// old terminal state
	oldTerminalState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	return oldTerminalState
}

func main() {
	oldTerminalState := enableRawMode()
	defer term.Restore(int(os.Stdin.Fd()), oldTerminalState)

	// we create a slice that stores 1 byte
	buf := make([]byte, 3)
	for {
		buf = make([]byte, 3)
		// we read the byte sotred in the slice
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}

		// cheking for arrow keys
		if n == 3 && buf[0] == 27 && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				fmt.Printf("Up arrow\r\n")
			case 'B':
				fmt.Printf("Down arrow\r\n")
			case 'C':
				fmt.Printf("Right arrow\r\n")
			case 'D':
				fmt.Printf("Left arrow\r\n")
			}
			continue
		}

		// if the user types CTRL+Q it quits.
		// 17 because in ASCII Q + 81 and CTRL removes 64 so 81-64 = 17

		const ctrlq = 17

		if buf[0] == ctrlq {
			break
		}
		// print other har (,only id n == 1)
		if n == 1 {
			if buf[0] < 32 || buf[0] == 127 {
				fmt.Printf("%d\r\n", buf[0])
			} else {
				fmt.Printf("%d ('%c')\r\n", buf[0], buf[0])
			}
		}
	}
}
