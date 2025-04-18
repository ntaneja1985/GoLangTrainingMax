package main

import "fmt"

// typr aliases
type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}
func main() {
	//maps.MainMap()

	//make functions take the datatype, length and capacity of the array
	//capacity is the maximum number of array items and determines the amount of memory space that will be allocated
	userNames := make([]string, 2, 5)

	userNames[0] = "William Shakespeare"
	userNames[1] = "John"
	//userNames = append(userNames, "Jane")
	//userNames = append(userNames, "Max")
	//userNames = append(userNames, "Max")
	//userNames = append(userNames, "Max")
	//userNames = append(userNames, "Max")
	//userNames = append(userNames, "Max")
	//fmt.Println(userNames)
	//fmt.Println(len(userNames))
	//fmt.Println(cap(userNames))
	//userNames = append(userNames, "Max")
	//userNames = append(userNames, "Max")
	//userNames = append(userNames, "Max")
	//fmt.Println(len(userNames))
	//fmt.Println(cap(userNames))

	//For maps, we can use make function and just specify the length of the map so that Go can pre-allocate memory
	//This makes it a bit more efficient
	courseRatings := make(floatMap, 5)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings.output()

	//If we dont care about individual item values or indexes, we can also just write for range userNames
	for index, value := range userNames {
		//fmt.Println(index)
		//fmt.Println(value)
		fmt.Println(userNames[index])
		fmt.Println(value)
		fmt.Println(value == userNames[index])
	}

	for key := range courseRatings {
		//fmt.Println(key)
		//fmt.Println(value)
		fmt.Println(courseRatings[key])
	}
}
