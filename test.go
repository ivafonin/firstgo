package main

/*go calculator*/

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	var num1, num2 int
	fmt.Println("Введите число 1 и 2:")
	fmt.Scanln(&num1, &num2)
	fmt.Println(sum(num1, num2))

}
