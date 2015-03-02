// golanghw project main.go
package main

import (
	"fmt"
)

func main() {
	var first_int int

	fmt.Println("Enter your integer:")
	fmt.Scanf("%d", &first_int)

	first_int = first_int / 2

	if first_int%2 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
