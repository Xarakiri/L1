package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования). */

// Human Родительская структура
type Human struct {
	Name   string
	Weight float32
}

// Eat Метод структуры Human
func (h *Human) Eat(calories float32) {
	h.Weight += calories / 1000
}

// Action Встраиваем методы от родительской структуры Human
type Action struct {
	Title string
	Human
}

func main() {
	action := Action{
		Title: "action",
		Human: Human{
			Name:   "John",
			Weight: 60,
		}}
	fmt.Println(action)

	action.Human.Eat(10000)
	fmt.Println(action)
}
