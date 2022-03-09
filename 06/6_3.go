package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
contex.WithCancel()
*/

func printer(ctx context.Context, ping, pong chan string) {
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

	// .WithCancel() завершает работу после вызова функции cancel()
	ctx, cancel := context.WithCancel(context.Background())

	ping <- "ping"
	go printer(ctx, ping, pong)

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("Программа завершилась после вызова функции cancel()!")
}
