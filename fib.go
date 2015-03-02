// golanghw project main.go
package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return n - 1 + n - 2
	}
}
func main() {
	var sum int
	var fnum int
	fmt.Println("Enter your n")
	fmt.Scanf("%d", &fnum)

	sum = fib(fnum)
	fmt.Println("The value is ", sum)
}
