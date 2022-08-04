package main

import "fmt"

func main() {
	//1. Слайсы (они же - срезы)
	// Слайс - это динамическая обвязка над массивом.
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] // Слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]:", startSlice)
	// Создали слайс, основываясь уже на существующем массиве

	//2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println("SecondSlice:", secondSlice)

	//3. Измнение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4] // Это набор ссылок на элементы нижележащего массива
	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

	//4. Один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modifications: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)

	//5. Срез как встроенный тип
	// type slice struct {
	// 	Length int
	// 	Capacity int
	// 	ZeroElement *byte
	// }

	//6. Длина и емкость слайса
	wordsSilce := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce))
	wordsSilce = append(wordsSilce, "four")
	fmt.Println("slice:", wordsSilce, "Length:", len(wordsSilce), "Capacity:", cap(wordsSilce))
	// Capacity (cap) или ёмкость слайса - это значение, показывающее СКОЛЬКО ЭЛЕМЕНТОВ В ПРИНЦИПЕ
	// можно добавить в срез БЕЗ ВЫДЕЛЕНИЯ ДОПОЛНИТЕЛЬНОЙ ПАМЯТИ ПОД НИЖЕЛЕЖАЩИЙ МАССИВ.

	// Допустим у нас есть срез на 3 элемента (инициализировали без явного указания массива)
	// Компилятор при создании этого среза СНАЧАЛА создал массив ровно на 3 элемента
	// После этого компилятор вернул адрес, где этот масив живет
	// Срез запомнил этот адрес и теперь ссылается на него
	// Потом
	// Начинаем деформировать слайс (увеличим длину /увеличим количество элементов)
	// Проблема - в нижележащем массиве (на котором основан слайс) все 3 ячейки. Что делать?
	// Компилятор ищет в памяти место для массива размера 3*2 (в общем случае n*2, где n - первоначальный размер)
	// После того как место найдено (в нашем случае найдено место для 6 элементов), в это место копируются
	// старые 3 элемента на свои позиции. На 4-ую позицию мы добавляем новый элемент
	// После этого компилятор возвращает нашему слайсу новый адрес в памяти, где находится массив под 6 элементов.

	//Емкость всегда будет изменять как n*2.
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	//Важно: после выделения памяти под новый массив, ссылки со старым будут перенсены в новый
	// Пример
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)

	// 7. Как создавать слайсы наиболее эффективно?
	// make() - это функция, позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)
	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	//Сначала инициализируется arr = [15]int
	//Затем по нему делается срез arr[0:10]
	//После чего он возаращается
	fmt.Println(sl)
	// 8. Добавление элементов в СРЕЗ
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherSlice := []string{"four", "five", "six"}
	myWords = append(myWords, anotherSlice...)
	myWords = append(myWords, "seven", "eight")
	fmt.Println("myWords:", myWords)
	//9. Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)

}
