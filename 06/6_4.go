package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
(использование отдельного канала))
*/

func printer(quit chan bool, ping, pong chan string) {
	for {
		select {
		case <-quit:
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
	// Канал сигнализирующий о закрытии горутины
	var quit chan bool = make(chan bool)
	defer close(ping)
	defer close(pong)

	ping <- "ping"
	go printer(quit, ping, pong)

	time.Sleep(time.Second * 3)
	// Посылаем сигнал для завершения горутины
	quit <- true
	fmt.Println("Программа завершилась после получения сигнала о завершении!")
}
