// write go a program to check whether a person is egiable to vote

package main

import "fmt"

func main() {
	var age int = 18

	if age >= 18 {
		fmt.Println("you are eligible to vote", age)
	} else {
		fmt.Println("you are not eligible to vote", age)
	}
}
