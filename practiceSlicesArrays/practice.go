package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	hobbies := [3]string{"Cricket", "Coding", "Azure"}
	fmt.Println("Hobbies: ", hobbies)
	fmt.Println("First element of Hobbies: ", hobbies[0])
	fmt.Println("Second and third element of Hobbies: ", hobbies[1:3])
	fmt.Println("First and third element of Hobbies: ", [2]string{hobbies[0], hobbies[2]})
	hobbiesSlice := hobbies[0:2]
	hobbiesSliceAnotherWay := hobbies[:2]
	fmt.Println(hobbiesSlice)
	fmt.Println(hobbiesSliceAnotherWay)
	hobbiesSlice = hobbies[1:]
	fmt.Println(hobbiesSlice)

	goals := []string{"Father", "Solution Architect"}
	fmt.Println(goals)
	goals[1] = "Azure Solution Architect"
	fmt.Println(goals)
	goals = append(goals, "Excellent Father")
	fmt.Println(goals)

	products := []Product{
		{title: "Simple Product 1", id: "1", price: 10},
		{title: "Simple Product 2", id: "2", price: 20},
	}
	fmt.Println(products)
	products = append(products, Product{
		title: "Simple Product 3",
		id:    "3",
		price: 30,
	})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
