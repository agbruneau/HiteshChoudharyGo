package main

import "fmt"

func main() {
	fmt.Println("Welcome to function inf golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Adder Result is:", result)

	proResult, proText := proAdder(2, 5, 8, 7, 3)

	fmt.Println(proText, proResult)
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Pro Added Result is:"
}
