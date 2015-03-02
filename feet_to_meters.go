// golanghw project main.go
package main

import (
	"fmt"
)

func main() {
	
	var feet int
	var meters int
	fmt.Println("Enter your feet:")
	fmt.Scanf("%d", &feet)
	meters = feet * 381 / 1250
	fmt.Println("It is ", meters, "meters.")
}
