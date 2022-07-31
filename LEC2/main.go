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

	//инициализация переменных
	var height int = 135
	fmt.Println(height)

	//инициализация скопа
	var (
		personName string = "Bob"
		personAge         = 42
		personUID  int
	)
	fmt.Printf("%s is %d years old with uid %d", personName, personAge, personUID)

	//странные штуки
	var a, b int = 30, 49
	fmt.Println(a, b)

	var c, d = 40, "String"
	fmt.Println(c, d)

	count := 10
	fmt.Println(count)
	count++
	fmt.Println(count)

	//множественное присваивание через :=

	e, y := 10, "vova"
	fmt.Println(e, y)
}
