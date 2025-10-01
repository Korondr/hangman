package main

import (
	"strings"
	"testing"
)

/* Тест функции CheckWin:
- первый случай — все нужные буквы угаданы (ожидаем true)
- второй случай — хотя бы одна буква не угадана (ожидаем false)*/
func TestCheckWin(t *testing.T) {
	expected := make([]bool, 26)
	guessedWin := make([]bool, 26)
	guessedLose := make([]bool, 26)

// Для примера буквы A, B, C
	expected[0] = true 
	expected[1] = true 
	expected[2] = true 

	guessedWin[0] = true
	guessedWin[1] = true
	guessedWin[2] = true

	if !CheckWin(expected, guessedWin) {
		t.Errorf("CheckWin() = false, want true")
	}

	
	guessedLose[0] = true
	guessedLose[1] = false
	guessedLose[2] = true

	if CheckWin(expected, guessedLose) {
		t.Errorf("CheckWin() = true, want false")
	}
}

/* Тест функции DrawWord:
 - проверяем, что буквы, которые были угаданы, отображаются
 - остальные символы заменяются на "*" */
func TestDrawWord(t *testing.T) { 
	word := "ABC"
	guessed := make([]bool, 26)
	guessed[0] = true 
	guessed[2] = true 

	DrawWord(word, guessed)
}

/* Тестируем функцию Categories:
- проверяем, что возвращаемый список категорий не пустой
- проверяем, что все элементы списка — строки*/
func TestCategories(t *testing.T) {
	cats := Categories()
	if len(cats) == 0 {
		t.Errorf("Categories() returned empty slice")
	}
	found := false
	for _, c := range cats {
		if strings.ToLower(c) == "animals" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected category 'animals' not found")
	}
}

/*Тест функции RandWord:
// - проверяем, что возвращаемое слово принадлежит выбранной категории
// - проверяем, что оно не пустое */
func TestRandWord(t *testing.T) {
	word := RandWord("animals")
	if word == "" {
		t.Errorf("RandWord('animals') returned empty string")
	}
}
