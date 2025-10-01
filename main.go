package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) >= 2 {
		// Неинтерактивный режим: go run . SECRET GUESSED
		RunNonInteractive(args[0], args[1])
	} else {
		// Интерактивный режим: go run .
		RunInteractive()
	}
}

