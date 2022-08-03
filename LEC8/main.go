package main

import "fmt"

func main() {
	//Массивы. Основа
	//1. Определение массива.
	//Создадим массив под хранение 5-ти целочисленных элементов
	var arr [5]int // При инициализации массива важно передать информацию -сколько элементов в нем будет
	fmt.Println("This is my array:", arr)
	//2. Определение элементов массива (после предварительной инициализации)
	// Необходимо обратиться к элементу массива через синтаксис arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[3] = -500
	arr[4] = 720
	fmt.Println("After elemtns init:", arr)
	//3. Определние массива с указанием элементов на месте
	// Если при инициализации количество элементов меньше номинальной длины массива
	// то недостающие элементы инициализируются нулями
	newArr := [5]int{10, 20, 30}
	fmt.Println("Short declaration and init:", newArr)
	//4. Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))
	arrWithValues[0] = 10000
	fmt.Println("Arr declaration with [...]:", arrWithValues, "Length:", len(arrWithValues))
	//5. Массив - это набор ЗНАЧЕНИЙ. То есть при всех манипуляциях - массив копируется (жестко, на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First arr:", first)
	fmt.Println("Second arr:", second)
	//6. Массив и его размер - это две составляющие одного типа (Размер массив - часть типа)
	// var aArr [5]int
	// var bArr [6]int
	// aArr[0] = 100
	// bArr = aArr
	// 7. Итерирование по массиву
	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}
	// 8. Итерирование по массиву через оператор range
	var sum float64
	for id, val := range floatArr {
		fmt.Printf("%d element of arr is %.2f\n", id, val)
		sum += val
	}
	fmt.Println("Total sum is:", sum)
	// 9. Игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}
	// 10. Многомерные массивы
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor", "Jo"},
	}
	fmt.Println("Multidimensional array:", words)
	// 11. Итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}

	//12. Общение со срезом == как массив без указания начального размера
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	fmt.Println(slice)
	slice = append(slice, 200)
	for i, v := range slice {
		fmt.Println(i, v)
	}
	emptySlice := []int{}
	emptySlice = append(emptySlice, 200)

}
