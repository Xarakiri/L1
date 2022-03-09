package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersect(set1, set2 map[string]bool) map[string]bool {
	answer := make(map[string]bool)
	for v := range set1 {
		_, ok := set2[v]
		if ok {
			answer[v] = true
		}
	}
	return answer
}

func main() {
	// Первое множество
	set1 := map[string]bool{"A": true, "B": true, "C": true}
	//Второе множество
	set2 := map[string]bool{"B": true, "C": true, "D": true}

	fmt.Println("Пересечение множеств: ", intersect(set1, set2))
}
