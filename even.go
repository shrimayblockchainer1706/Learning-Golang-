// Write a go program to check whether the given interger is even or odd

package main

import "fmt"

func main() {
	a := 10

	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
