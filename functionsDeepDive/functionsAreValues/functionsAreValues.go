package functionsAreValues

import "fmt"

type transformFn func(int) int

func functionsAreValues() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moreNumbers := []int{11, 12, 13, 14}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunctionWithNumbers(&numbers)
	transformerFn2 := getTransformerFunctionWithNumbers(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func getTransformerFunction() transformFn {
	return double
}

func getTransformerFunctionWithNumbers(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func triple(number int) int {
	return number * 3
}
