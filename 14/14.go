package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel
из переменной типа interface{}.
*/

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int!\n", i)
	case string:
		fmt.Printf("%v is string!\n", i)
	case bool:
		fmt.Printf("%v is bool!\n", i)
	case chan int:
		fmt.Printf("%v is chan int!\n", i)
	case chan bool:
		fmt.Printf("%v is chan bool!\n", i)
	case chan string:
		fmt.Printf("%v is chan string!\n", i)
	default:
		fmt.Printf("I cant define type of %v!\n", v)
	}
}

func main() {
	var i interface{}
	i = 1

	// 1й способ с помощью форматирования строки
	fmt.Printf("Type of i is %T\n", i)

	// 2й способ с помощью switch
	printType(i)

	// 3й способ с помощью пакета reflect
	fmt.Println(reflect.TypeOf(i))
}
