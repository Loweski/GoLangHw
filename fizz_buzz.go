// golanghw project main.go
package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 100 {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i = i + 1
	}
}
