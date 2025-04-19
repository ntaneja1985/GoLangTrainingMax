package anonymous

import (
	"fmt"
)

func Anonymous() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	fmt.Println(transformed)

	double := createTransformer(2)
	triple := createTransformer(3)
	fmt.Println(double(4))
	fmt.Println(triple(5))

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
