package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}

// Multiply two numbers
func Multiply(number1, number2 int) int {
	return number1 * number2
}

// Divide two numbers
func Divide(number1, number2 float32) float32 {
	return number1 / number2
}

// Subtract two numbers
func Subtract(number1, number2 int) int {
	return number1 - number2
}
