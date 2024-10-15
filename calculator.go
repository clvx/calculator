package calculator

// create a function to calculate the sum of two numbers
func add(x int, y int) int {
	return x + y
}

// create a function to calculate the difference of two numbers
func subtract(x int, y int) int {
	return x - y
}

// create a function to calculate the product of two numbers
func multiply(x int, y int) int {
	return x * y
}

// create a function to divide two numbers
func divide(x int, y int) (int, error) {
    if y == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return x / y, nil
}
