package main

func main() {
	Run()
}

func Run() []int {
	results := []int{
		sum(2, 2),
		subtract(5, 3),
		multiply(4, 2),
		sumX(2, 2),
	}
	return results
}

func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func sumX(a int, b int) int {
	return a + b + a
}
