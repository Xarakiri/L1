package main

import (
	"fmt"
	"unicode"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func IsUnique(s string) bool {
	m := make(map[rune]bool)

	for _, v := range s {
		if _, ok := m[unicode.ToLower(v)]; ok {
			return false
		}
		m[v] = true
	}
	return true
}

func main() {
	input := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, v := range input {
		fmt.Println(IsUnique(v))
	}
}
