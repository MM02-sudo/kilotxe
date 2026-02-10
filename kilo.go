package main

import (
	"os"
)

func main() {
	buf := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		if buf[0] == 'q' {
			break
		}

	}
}
