package main

import (
	"fmt"
)

func main() {
	fmt.Println("======= Functions =======")

	a := 2
	b := 5
	var c int
	fmt.Printf("\nPass by value: %d", sumByValue(a, b))

	sumByRef(a, b, &c)
	fmt.Printf("\nPass by reference: %d", c)

	numbers := []int{2, 2, 4, 5, 12}
	fmt.Printf("\nSum of Numbers: %d", testingManyVars(numbers...))
}
func testingManyVars(sArr ...int) (total int) {
	for _, num := range sArr {
		total += num
	}
	return total
}

func sumByValue(a, b int) int {
	return a + b
}

func sumByRef(a, b int, c *int) {
	*c = a + b
}
