package main 

import(
	"fmt"
)


func main(){
	var numberOne int = 9
	var numberTwo int = 10

	result := add(numberOne, numberTwo)

	fmt.Println("Addition Result:",numberOne,"+",numberTwo,"=",result)

	subResult := sub(numberOne, numberTwo)

	fmt.Println("Subtraction Result:", numberOne, "-", numberTwo, "=", subResult)

	mResult := mult(numberOne, numberTwo)

	fmt.Println("Multiplication Result:", numberOne, "x", numberTwo, "=", mResult)

	dResult := divide(numberOne, numberTwo)

	fmt.Println("Division Result:", numberOne, "/", numberTwo, "=", dResult)
}

// Adds two numbers together
func add(numOne int, numTwo int) int{
	result := numOne + numTwo
	return result
}

// Subtract one number from another
func sub(numOne int, numTwo int) int {
	result := numOne - numTwo
	return result
}

// Multiply one number by the second
func mult(numOne int, numTwo int) int {
	return numOne * numTwo
}

// Divide one number by the second
func divide(numOne int , numTwo int) float32 {
	return float32(numOne) / float32(numTwo)
}
