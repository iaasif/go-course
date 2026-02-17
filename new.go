package main

import "fmt"

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

func test() {
	// ============================================
	// 1) Create a new array with 3 hobbies
	// ============================================
	hobbies := [3]string{"Reading", "Gaming", "Coding"}
	fmt.Println("1) Hobbies array:", hobbies)

	// ============================================
	// 2) Output more data about the array
	// ============================================
	// First element (standalone)
	fmt.Println("2a) First element:", hobbies[0])

	// Second and third element combined as a new list (slice)
	secondAndThird := hobbies[1:3]
	fmt.Println("2b) Second and third elements:", secondAndThird)

	// ============================================
	// 3) Create a slice with first and second elements
	//     Two different ways
	// ============================================
	// Way 1: Using slice expression [low:high]
	slice1 := hobbies[0:2]
	fmt.Println("3a) Slice way 1 (hobbies[0:2]):", slice1)

	// Way 2: Using [:high] shorthand (low defaults to 0)
	slice2 := hobbies[:2]
	fmt.Println("3b) Slice way 2 (hobbies[:2]):", slice2)

	// ============================================
	// 4) Re-slice the slice from (3) to contain
	//    the second and last element of original array
	// ============================================
	// Re-slicing slice1 to get second element (index 1) to end
	reSliced := slice1[1:2] // Gets "Gaming" which is 2nd element of original
	fmt.Println("4) Re-sliced (second element):", reSliced)

	// Alternative: Get second AND last element (both 2nd and 3rd from original)
	secondAndLast := hobbies[1:3] // This gets "Gaming" and "Coding"
	fmt.Println("4b) Second and last elements:", secondAndLast)

	// ============================================
	// 5) Create a "dynamic array" (slice) with course goals
	// ============================================
	courseGoals := []string{"Learn Go", "Build Projects"}
	fmt.Println("5) Course goals:", courseGoals)

	// ============================================
	// 6) Modify second goal AND add a third goal
	// ============================================
	// Modify second goal (index 1)
	courseGoals[1] = "Master Go"
	fmt.Println("6a) After modifying second goal:", courseGoals)

	// Add a third goal using append
	courseGoals = append(courseGoals, "Get Certified")
	fmt.Println("6b) After adding third goal:", courseGoals)

	// ============================================
	// 7) BONUS: Product struct and dynamic list
	// ============================================
	// Define Product struct
	type Product struct {
		title string
		id    int
		price float64
	}

	// Create a dynamic list (slice) of products
	products := []Product{
		{title: "Laptop", id: 1, price: 999.99},
		{title: "Mouse", id: 2, price: 29.99},
	}

	fmt.Println("7) Products:")
	for i, p := range products {
		fmt.Printf("   Product %d: Title=%s, ID=%d, Price=$%.2f\n", i+1, p.title, p.id, p.price)
	}

	// Add another product
	products = append(products, Product{title: "Keyboard", id: 3, price: 79.99})
	fmt.Println("7b) After adding product:")
	for i, p := range products {
		fmt.Printf("   Product %d: Title=%s, ID=%d, Price=$%.2f\n", i+1, p.title, p.id, p.price)
	}
}
