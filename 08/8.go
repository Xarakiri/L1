package main

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func reverseBit(v int64, bit, flag uint8) int64 {
	// Создаем маску которая содержит один бит в указанной позиции
	var mask int64 = 1 << bit

	// Если бит нужно поставить в 0
	// То реверсируем маску и приминяем побитовое И
	if flag == 0 {
		return v & ^mask
	}
	// Иначе применяем побитовое ИЛИ
	return v | mask
}

func main() {
	var v int64
	var bit uint8
	var flag uint8

	fmt.Print("Введите число int64: ")
	_, err := fmt.Scanln(&v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите номер бита i: ")
	_, err = fmt.Scanln(&bit)
	if err != nil {
		fmt.Println(err)
		return
	}
	if bit < 0 || bit > 63 {
		fmt.Println("Incorrect i!")
		return
	}


	fmt.Print("Введите значение бита 1 или 0: ")
	_, err = fmt.Scanln(&flag)
	if err != nil {
		fmt.Println(err)
		return
	}
	if flag != 0 && flag != 1 {
		fmt.Println("Incorrect flag!")
		return
	}

	fmt.Printf("Result: %d\n", reverseBit(v, bit, flag))
}