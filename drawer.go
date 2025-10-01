package main

import "fmt"

var hangmanStates = []string{
	`
	+---+
	    |
	    |
	    |
	   ===`,
	`
	+---+
	O   |
	    |
	    |
	   ===`,
	`
	+---+
	O   |
	|   |
	    |
	   ===`,
	`
	+---+
	  O   |
   /|   |
	      |
	   ===`,
	`
	+---+
	  O   |
   /|\  |
	      |
	   ===`,
	`
	+---+
	  O   |
   /|\  |
   /    |
	   ===`,
	`
	+---+
	  O   |
   /|\  |
   / \  |
	   ===`,
}

func DrawHangman(mistakes int) {
	if mistakes >= len(hangmanStates) {
		mistakes = len(hangmanStates) - 1
	}
	fmt.Println(hangmanStates[mistakes])
}

