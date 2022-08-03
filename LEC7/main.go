package main

import "fmt"

func main() {
	//Switch!
	var price int
	fmt.Scan(&price)
	//В switch-case запрещены дублирующиеся условия в case"ах
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")
	default:
		//Отрабатывает только в том случае, если не один из выше перечисленных кейсов - не сработал
		fmt.Println("Default case")

	}

	//Case с множеством вариантов
	var ageGroup string = "q" //Возрастнаые группы : "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You are too yong/old")
	}

	// Case с выражениями
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("Too yong")
	case age > 18 && age <= 30:
		fmt.Println("Second case")
	case age > 30 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you")
	}

	//Case с проваливаниями. Проваливания выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ
	//В момент выполнения fallthroug у следующего кейса не проверяется условие,
	//а сразу выполняется тело
	var number int
	fmt.Scan(&number)
outer:
	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		if number%2 == 0 {
			break outer
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d GREATER then 200\n", number)
		fallthrough
	case number > 1000:
		fmt.Printf("%d GREATER then 1000\n", number)
		fallthrough
	default:
		fmt.Printf("%d DEFAULT\n", number)
	}

	//Гадость с терминацией цикла for из switchv
	var i int
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}

	fmt.Println("END")
}
