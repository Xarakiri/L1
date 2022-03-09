package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
contex.WithDeadline()
*/

func printer(ctx context.Context, wg *sync.WaitGroup, ping, pong chan string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ping:
			fmt.Println(msg)
			pong <- "pong"
			time.Sleep(time.Second)
		case msg := <-pong:
			fmt.Println(msg)
			ping <- "ping"
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создадим канал
	var ping chan string = make(chan string, 1)
	var pong chan string = make(chan string, 1)
	defer close(ping)
	defer close(pong)

	// .WithDeadline() завершает работу к определнному времени
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	wg := &sync.WaitGroup{}

	wg.Add(1)
	ping <- "ping"
	go printer(ctx, wg, ping, pong)

	wg.Wait()
	cancel()
	fmt.Println("Программа завершилась time.Now()+3 секунды!")
}
