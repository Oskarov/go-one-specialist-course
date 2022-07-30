package main

import "fmt"

func main() {
	fmt.Println("Простейший вывод в консоль с переносом строки")
	fmt.Print("Вывод аргумента")
	fmt.Printf("стандартный вывод с флагами форматирования - %s\n", "Дима")

	//декларирование переменных

	var age int
	age = 32
	fmt.Println(age)

	var height int = 135
	fmt.Println(height)

}
