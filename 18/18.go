package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// Создаем нашу структуру
	c := Counter{
		mu:    sync.Mutex{},
		count: 0,
	}
	wg := sync.WaitGroup{}

	// Запускаем 1000 горутин, которые инкрементируют счетчик
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	// Ждем завершения работ горутин
	wg.Wait()

	fmt.Println("Counter:", c.Get())
}
