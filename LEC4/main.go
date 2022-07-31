package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//boolean по умолчанию false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	aBoolean, bBoolean := true, false
	fmt.Println("and:", aBoolean && bBoolean)
	fmt.Println("or:", aBoolean || bBoolean)
	fmt.Println("not:", !aBoolean)

	//numbers. Integers
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, uint64, uint
	var aNum int = 32
	bNum := 92
	fmt.Println("sum = ", aNum+bNum)

	//Вывод типа
	fmt.Printf("type is %T", aNum)
	//Узнаем сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes", aNum, unsafe.Sizeof(aNum))

	//Явное приведение типов при необходимости
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(first32 + int32(second64))
	//если уверены, что не произойдёт коллизии, в противном случае часть значения не влезет
	// + - * / %

	//float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of %T", floatFirst, floatSecond)
	fmt.Printf("sum: %.3f and Sub: %.3f", floatFirst+floatSecond, floatFirst-floatSecond)

	//numeric. complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

}
