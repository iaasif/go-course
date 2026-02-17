package main

func main() {

}

func add[T int | float64 | string](a, b T) T {
	return a + b

}
