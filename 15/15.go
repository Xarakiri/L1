package main

/*Задание:
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}
*/

// В функции мы генерим строку длиной 1024 символа
// затем глобальная переменная justString ссылается на локальную переменную v
// так как на v весит ссылка, то Garbage Collector не почистит v
// вследсвии чего v занимает лишнюю память.
// Решить проблему можно просто создав новую строку из слайса:
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = string(v[:100])
//}
// При этом выделится новая область памяти в куче и на v уже ничего ссылаться не будет.