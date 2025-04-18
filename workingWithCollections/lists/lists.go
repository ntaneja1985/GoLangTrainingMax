package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main1() {
	//Go automatically creates a slice for us and it also creates an underlying array
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[1])
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	//Throws errors as we cannot access indexes that dont exist
	//prices[2] = 11.99

	//Append function not only adds the element to the slice but also to the underlying array
	updatedPrices := append(prices, 14.99)
	//Note the original slice doesnot change, to do that we have to do prices = append(prices,14.99)
	//Also note that the underlying array is automatically updated by Go
	//It is typical to work with slices in Go rather than use arrays
	//If we are sure about the number of the elements in the list, go for arrays
	fmt.Println(updatedPrices, prices)

	//Remove the elements
	//No function as such to remove elements since we already have slices
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}

//func main() {
//	var productNames [4]string = [4]string{"1", "2", "3"}
//	productNames[3] = "4"
//	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
//	//fmt.Println(productNames)
//	//fmt.Println(prices[2])
//
//	//Start with the first element at the index 1(include element at index 1)
//	//and go upto index 3 and exclude element at index 3, so this will have only 2 values
//	//featuredPrices := prices[1:3]
//	//fmt.Println(featuredPrices)
//
//	//We can omit the starting index and write the ending index
//	//This will start from the first index i.e 0 and go all the way upto index 3
//	//and exclude the element at index 3
//	//newPrices := prices[:3]
//	//fmt.Println(newPrices)
//
//	//Can do opposite also
//	//Start with element at index 1 and go all the way upto the end element
//	//testPrices := prices[1:]
//	//fmt.Println(testPrices)
//
//	//Using highest bound in a slice
//	//Go has a special feature where we can go upto lastIndex + 1
//	//This allows us to include the last element in the slice also(remember element at ending index is excluded)
//	//includeLastElementPrices := prices[1:4]
//	//fmt.Println(includeLastElementPrices)
//
//	//Slices can be used on arrays, but they can also be based on other slices
//	fmt.Println("Original Prices array: ", prices) //Returns [10.99 9.99 45.99 20]
//	firstSlice := prices[1:]                       //Returns [9.99, 45.99, 20.0]
//	fmt.Println("First Slice: ", firstSlice)
//	secondSlice := firstSlice[:1] //Returns [9.99]
//	fmt.Println("Second Slice: ", secondSlice)
//	secondSlice[0] = 17.99
//	fmt.Println("Second Slice after modification: ", secondSlice)
//	fmt.Println("Original Prices array second time after modification: ", prices)
//	fmt.Println("Length of first slice: ", len(firstSlice))   //Gives 3
//	fmt.Println("Capacity of first slice: ", cap(firstSlice)) //Gives 3
//	fmt.Println("Length of second slice: ", len(secondSlice))
//	fmt.Println("Capacity of second slice: ", cap(secondSlice))
//
//	randomFirstSlice := prices[0:4]
//	fmt.Println("Random First Slice: ", randomFirstSlice)
//	fmt.Println("Length of random first slice: ", len(randomFirstSlice))
//	fmt.Println("Capacity of random first slice: ", cap(randomFirstSlice))
//	randomSecondSlice := randomFirstSlice[1:]
//	fmt.Println("Random Second Slice: ", randomSecondSlice)
//	fmt.Println("Length of random second slice: ", len(randomSecondSlice))
//	fmt.Println("Capacity of random second slice: ", cap(randomSecondSlice))
//	randomThirdSlice := randomFirstSlice[:1]
//	fmt.Println("Random Third Slice: ", randomThirdSlice)
//	fmt.Println("Length of random third slice: ", len(randomThirdSlice))
//	fmt.Println("Capacity of random third slice: ", cap(randomThirdSlice))
//
//	randomThirdSlice = randomThirdSlice[1:]
//	fmt.Println("Random Third Slice: ", randomThirdSlice)
//	fmt.Println("Length of random third slice: ", len(randomThirdSlice))
//	fmt.Println("Capacity of random third slice: ", cap(randomThirdSlice))
//}
