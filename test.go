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
	var choise string
	fmt.Println("Введите операцию : +/-/:/*")
	fmt.Scanln(&choise)
	switch choise {
	case "+":
		fmt.Println(num1, " + ", num2, " = ", sum(num1, num2))
	case "*":
		fmt.Println(num1, " * ", num2, " = ", num1*num2)
	case ":":
		fmt.Println(num1, " : ", num2, " = ", num1/num2)
	case "-":
		fmt.Println(num1, " - ", num2, " = ", num1-num2)
	default:
		fmt.Println("Unavalibale")
	}

}
