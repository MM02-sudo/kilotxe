package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func enableRawMode() *term.State {
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
	buf := make([]byte, 1)
	for {
		// we read the byte sotred in the slice
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		// if the user types CTRL+Q it quits.
		// 17 because in ASCII Q + 81 and CTRL removes 64 so 81-64 = 17

		const ctrlq = 17

		if buf[0] == ctrlq {
			break
		}
		// writing back to screen each key stroke of the user
		fmt.Printf("%d\r\n", buf[0])

	}
}
