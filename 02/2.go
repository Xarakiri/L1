package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и
выведет их квадраты в stdout.
*/

func square(number int) int {
	return number * number
}

func main() {
	input := [...]int{2, 4, 6, 8, 10}

	// Реализация с использованием WaitGroup
	fmt.Println("Wait groups:")
	wg := sync.WaitGroup{}
	for _, v := range input {
		wg.Add(1)
		go func(compute int) {
			defer wg.Done()
			fmt.Println(square(compute))
		}(v)
	}
	wg.Wait()

	// Реализация с использованием channel
	fmt.Println("Channels:")
	// Создаем буфферизированный канал, для синхронизации горутин
	ch := make(chan int)
	defer close(ch)
	for _, v := range input {
		go func(compute int) {
			ch <- square(compute)
		}(v)
	}

	// Функция для получения данных с канала и вывода их в stdout
	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
		}
	}()
	// Ждем ввода с консоли чтобы не вылетел main
	fmt.Scanln()
}
