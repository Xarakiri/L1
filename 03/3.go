package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
func waitgroup() {
	input := [...]int{2, 4, 6, 8, 10}

	answer := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// Запускаем горутины
	for _, v := range input {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()

			// Блокируем доступ к answer
			mu.Lock()
			answer += value * value
			mu.Unlock()
		}(v)
	}

	// Ждем завершения горутин
	wg.Wait()

	fmt.Printf("WaitGroup Sum: %v\n", answer)
}

func channel() {
	var answer int
	input := []int{2, 4, 6, 8, 10}

	// Создаем канал
	ch := make(chan int, len(input))
	defer close(ch)

	// Пишем квадраты чисел в канал
	for _, v := range input {
		go func(v int, c chan int) {
			c <- v * v
		}(v, ch)
	}

	// Читаем данные из канала
	for i := 0; i < len(input); i++ {
		answer += <-ch
	}

	// Выводим ответ в консоль
	fmt.Printf("Channels Sum: %v\n", answer)
}

func main() {
	waitgroup()
	channel()
}
