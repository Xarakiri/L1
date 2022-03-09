package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/
func genArr(len int) []int {
	r := make([]int, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		r[i] = rand.Intn(100)
	}
	return r
}

func binSearch(arr []int, toSearch int) int {
	left := 0
	right := len(arr)

	for left <= right {
		mid := (left + right) / 2
		if toSearch < arr[mid] {
			right = mid - 1
			continue
		}
		if toSearch > arr[mid] {
			left = mid + 1
			continue
		}
		if toSearch == arr[mid] {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{5, 14, -1, 3, 22, 17, 44, 51, 20, 3, 19}
	toSearch := 17

	// Сортируем
	sort.Ints(arr)

	fmt.Println("Sorted array:", arr)
	fmt.Printf("Index of %d = %d\n", toSearch, binSearch(arr, toSearch))
}
