package main

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func reverse(s string) string {
	// Делаем из строки, слайс рун (для работы с юникод)
	r := []rune(s)

	l := len(r)

	// Переворачиваем строку
	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}
	return string(r)
}

func main() {
	input := "🤒🤕🤮"

	fmt.Println("Reversed string:", reverse(input))
}
