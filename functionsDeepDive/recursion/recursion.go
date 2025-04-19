package recursion

import "fmt"

func Recursion() {
	fact := factorial(5)
	fmt.Println(fact)
}

//func factorial(num int) int {
//	result := 1
//	for i := 1; i <= num; i++ {
//		result *= i
//	}
//	return result
//}

func factorial(num int) int {
	//exit condition
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
