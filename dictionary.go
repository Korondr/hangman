package main

import "math/rand"

var wordCategories = map[string][]string{
	"animals": {"elephant", "tiger", "cat", "dog"},
	"fruits":  {"apple", "banana", "orange"},
}

func Categories() []string {
	cats := make([]string, 0, len(wordCategories))
	for k := range wordCategories {
		cats = append(cats, k)
	}
	return cats
}

func RandWord(category string) string {
	words, ok := wordCategories[category]
	if !ok || len(words) == 0 {
		// случайная категория
		for _, w := range wordCategories {
			words = w
			break
		}
	}
	return words[rand.Intn(len(words))]
}

