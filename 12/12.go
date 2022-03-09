package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func makeSet(seq []string) map[string]bool {
	answer := make(map[string]bool)

	// Так как ключи в словаре уникальны, у нас получится собственное множество
	for _, v := range seq {
		answer[v] = true
	}

	return answer
}

func main() {
	// Последовательность строк
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("Собственное множество:", makeSet(seq))
}
