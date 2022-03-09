package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func Sleep1(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}

func Sleep2(sec int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	defer cancel()
	<-ctx.Done()
}

func main() {
	Sleep1(1)
	Sleep2(1)
	fmt.Println("Программа завершилась спустя 2 секунды!")
}
