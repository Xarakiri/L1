package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverse(s string) string {
	// Сплитим строку по пробелу
	r := strings.Split(s, " ")

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return strings.Join(r, " ")
}

func main() {
	input := "snow dog sun"

	fmt.Println("Reversed string:", reverse(input))
}
