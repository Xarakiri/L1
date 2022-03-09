package main

import (
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	// Инициализируем мапу
	m := make(map[int]int)

	// Можно использовать sync.Map{}
	// и тогда можно не блокировать руками доступ на чтение/запись
	// ma := sync.Map{}
	// ma.Store(1, 1)

	// Создаем мьютекст
	mu := &sync.Mutex{}
	// Создаем вейтгруппу
	wg := &sync.WaitGroup{}

	// Запускаем горутины
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			// Блокируем доступ на запись к мапе
			// Если этого не делать, то возникнет ошибка:
			// fatal error: concurrent map writes
			mu.Lock()
			m[1] = v
			mu.Unlock()
		}(i)
	}
	// Ждем пока горутины завершат работу
	wg.Wait()
}
