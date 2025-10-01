package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	easyAttem   = 8
	mediumAttem = 6
	hardAttem   = 4
)

func RunInteractive() {
	fmt.Println("Welcome to Hangman")

	difficulty := ChooseDifficulty()
	attempts := AttemptsByDifficulty(difficulty)
	categories := Categories()
	fmt.Println("Available categories:", categories)
	category := ChooseCategory(categories)
	word := RandWord(category)

	word = strings.ToUpper(word)
	ExpectedLetters := make([]bool, 26)
	GuessedLetters := make([]bool, 26)

	for _, ch := range word {
		if unicode.IsLetter(ch) {
			ExpectedLetters[ch-'A'] = true
		}
	}

	mistakes := 0
	for {
		DrawHangman(mistakes)
		DrawWord(word, GuessedLetters)
		fmt.Printf("\nAttempts left: %d\n", attempts-mistakes)

		if mistakes >= attempts {
			fmt.Println("\nYou lost! The word was:", word)
			break
		}
		if CheckWin(ExpectedLetters, GuessedLetters) {
			fmt.Println("\nYou won!")
			break
		}

		var input string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&input)

		// Проверка корректности ввода: только один символ
		if len(input) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		letter := unicode.ToUpper(rune(input[0]))

		// Проверка, что введена допустимая буква (A-Z)
		if letter < 'A' || letter > 'Z' {
			fmt.Println("Invalid input. Only letters are allowed.")
			continue
		}

		index := letter - 'A'

		// Проверка, была ли эта буква уже угадана
		if GuessedLetters[index] {
			fmt.Println("You already guessed this letter.")
			continue
		}

		GuessedLetters[index] = true
		
		// Если буквы нет в слове — увеличиваем число ошибок
		if !ExpectedLetters[index] {
			mistakes++
			fmt.Println("Wrong guess!")
		} else {
			fmt.Println("Correct guess!")
		}
	}
}

func RunNonInteractive(secret, guessed string) {
	secret = strings.ToUpper(secret)
	guessed = strings.ToUpper(guessed)

	ExpectedLetters := make([]bool, 26)
	GuessedLetters := make([]bool, 26)

	for _, ch := range secret {
		if unicode.IsLetter(ch) {
			ExpectedLetters[ch-'A'] = true
		}
	}
	for _, ch := range guessed {
		if unicode.IsLetter(ch) {
			GuessedLetters[ch-'A'] = true
		}
	}

	result := ""
	win := CheckWin(ExpectedLetters, GuessedLetters)
	for _, ch := range secret {
		if GuessedLetters[ch-'A'] {
			result += string(ch)
		} else {
			result += "*"
		}
	}

	status := "lose"
	if win {
		status = "win"
	}

	fmt.Printf("%s;%s\n", result, status)
}

func ChooseDifficulty() string {
	var diff string
	for {
		fmt.Println("Choose difficulty (easy, medium, hard):")
		fmt.Scanln(&diff)
		diff = strings.ToLower(diff)
		if diff == "easy" || diff == "medium" || diff == "hard" {
			return diff
		}
		fmt.Println("Invalid difficulty, try again.")
	}
}

func AttemptsByDifficulty(diff string) int {
	switch diff {
	case "easy":
		return easyAttem
	case "medium":
		return mediumAttem
	case "hard":
		return hardAttem
	default:
		return mediumAttem
	}
}

func ChooseCategory(categories []string) string {
	fmt.Println("Available categories:")
	for _, c := range categories {
		fmt.Printf("- %s\n", c)
	}

	valid := make(map[string]bool)
	for _, c := range categories {
		valid[strings.ToLower(c)] = true
	}

	var choice string
	for {
		fmt.Println("Choose category (or press Enter for random):")
		fmt.Scanln(&choice)
		choice = strings.ToLower(choice)

		if choice == "" || valid[choice] {
			return choice
		}
		fmt.Println("Invalid category, try again.")
	}
}

func CheckWin(expected, guessed []bool) bool {
	for i := 0; i < 26; i++ {
		if expected[i] && !guessed[i] {
			return false
		}
	}
	return true
}

func DrawWord(word string, guessed []bool) {
	for _, ch := range word {
		if guessed[ch-'A'] {
			fmt.Printf("%c ", ch)
		} else {
			fmt.Print("* ")
		}
	}
	fmt.Println()
}

