package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func deleteI(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		panic("Error!")
	}
	low := s[:i]
	high := s[i+1:]

	return append(low, high...)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(deleteI(s, 3))
}
