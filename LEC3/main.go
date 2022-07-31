package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)
	/*fmt.Scan(&age)
	fmt.Scan(&name)*/

	_, err := fmt.Scan(&age, &name)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	fmt.Printf("my name id: %s and my age is: %d", name, age)

}
