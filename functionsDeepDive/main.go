package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup("testing", 2, numbers...)
	fmt.Println(sum)
}
func sumup(prestarting string, startingValue int, numbers ...int) int {
	fmt.Println(prestarting)
	sum := startingValue
	for _, number := range numbers {
		sum += number
	}
	return sum
}
