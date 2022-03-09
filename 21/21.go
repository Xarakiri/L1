package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// У нас есть европейска розетка и китайская вилка.
// Чтобы их подружить, создаем новый класс Адаптер - переходник с китайской вилки на европейскую.

type PowerSocketEU struct{}

func (c *PowerSocketEU) InsertPlugIntoSocket(plug Plug) {
	plug.InsertPlugEU()
	fmt.Println("Вилка вошла в европейскую розетку.")
}

type Plug interface {
	InsertPlugEU()
}

type PlugCN struct{}

func (c *PlugCN) InsertPlugCN() {
	fmt.Println("Вилка вошла в китайскую розетку.")
}

type AdapterCN struct {
	plug *PlugCN
}

func (a *AdapterCN) InsertPlugEU() {
	a.plug.InsertPlugCN()
	fmt.Println("(Китайская вилка вошла в переходник)")
}

func main() {
	eu := &PowerSocketEU{}
	cn := &PlugCN{}
	adapter := &AdapterCN{plug: cn}

	eu.InsertPlugIntoSocket(adapter)
}
