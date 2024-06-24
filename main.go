package main

import (
	"fmt"
	"operations/addition"
	"operations/division"
	"operations/ispositive"
	"operations/multiply"
	"operations/subtract"
)

func main() {
	result1 := ispositive.IsPositive(5)
	fmt.Println(result1)
	result2 := addition.Add(5, 6)
	fmt.Println(result2)
	result3 := subtract.Subtract(7, 3)
	fmt.Println(result3)
	result4 := multiply.Multiply(3, 5)
	fmt.Println(result4)
	result5, err := division.Divide(15, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result5)
}
