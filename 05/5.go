package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func produce(ch chan<- int) {
	event := 0
	for {
		ch <- event
		event++
		time.Sleep(time.Millisecond * 500)
	}
}

func read(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		select {
		case job := <-ch:
			fmt.Printf("Job - %d\n", job)
		case <-ctx.Done():
			close(ch)
			return
		}
	}
}

func main() {
	var n int
	fmt.Print("Введите время (в секундах) через которое программа должна завершиться: ")
	fmt.Scanln(&n)

	// Создадим канал
	jobsChan := make(chan int)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go produce(jobsChan)
	go read(ctx, wg, jobsChan)

	wg.Wait()
	cancelFunc()
	fmt.Printf("Программа завершила свою работу по истечению %d секунд\n", n)
}
