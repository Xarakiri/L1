package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a, b := 777, 666

	// Меняем числа местами
	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)
}