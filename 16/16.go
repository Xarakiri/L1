package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func genArr(len int) []int {
	r := make([]int, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		r[i] = rand.Intn(100)
	}
	return r
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]

	low := make([]int, 0, len(arr))
	mid := make([]int, 0, len(arr))
	high := make([]int, 0, len(arr))

	for _, i := range arr {
		switch {
		case i < pivot:
			low = append(low, i)
		case i == pivot:
			mid = append(mid, i)
		case i > pivot:
			high = append(high, i)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	low = append(low, mid...)
	low = append(low, high...)

	return low
}

func main() {
	arr := genArr(10)
	fmt.Println("Before sort:", arr)
	fmt.Println("After sort:", quickSort(arr))
}
