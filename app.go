package main

import "fmt"

// Product struct for exercise 7
type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1) Create a new array with three hobbies and output it
	hobbies := [3]string{"reading", "gaming", "hiking"}
	fmt.Println("1) Hobbies array:", hobbies)

	// 2) Output first element standalone, and second+third as new list
	fmt.Println("   First element:", hobbies[0])
	slice2and3 := hobbies[1:3]
	fmt.Println("   Second and third:", slice2and3)

	// 3) Create two slices containing first and second elements (different ways)
	// Way 1: Slicing from array
	slice1 := hobbies[0:2]
	fmt.Println("3) Slice way 1:", slice1)

	// Way 2: Using literal declaration
	slice2 := []string{hobbies[0], hobbies[1]}
	fmt.Println("   Slice way 2:", slice2)

	// 4) Re-slice to contain second and last element
	reSlice := slice1[1:3] // or hobbies[1:3]
	fmt.Println("4) Re-sliced (2nd & last):", reSlice)

	// 5) Create dynamic array with course goals
	goals := []string{"Learn Go fundamentals", "Build a web API"}
	fmt.Println("5) Goals:", goals)

	// 6) Change second goal and add a third
	goals[1] = "Master concurrency patterns"
	goals = append(goals, "Deploy a microservice")
	fmt.Println("6) Updated goals:", goals)

	// 7) Bonus: Product struct with dynamic list
	products := []Product{
		{title: "Laptop", id: 1, price: 999.99},
		{title: "Mouse", id: 2, price: 29.99},
	}
	fmt.Println("7) Products:")
	for _, p := range products {
		fmt.Printf("   ID: %d, Title: %s, Price: $%.2f\n", p.id, p.title, p.price)
	}
}
