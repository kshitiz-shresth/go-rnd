package main

import "fmt"

func main() {
	var funcAsArgumentAndReturnFunc = func(addFunc func(int, int) int, arg1 int, arg2 int) func(int) int {
		var mySum = addFunc(arg1, arg2)
		return func(numArg int) int {
			return mySum + numArg
		}
	}

	var out = funcAsArgumentAndReturnFunc(
		func(a int, b int) int {
			return a + b
		},
		5,
		10,
	)(10)

	fmt.Println("Output as function as argument is", out)

	fmt.Println()
	var nested = func() {
		fmt.Println("top function")
		var nested2 = func() {
			fmt.Println("middle function")
			var nested3 = func() {
				fmt.Println("inner function")
			}
			nested3()
		}
		nested2()
	}
	nested()

}
