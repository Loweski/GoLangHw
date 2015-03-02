// golanghw project main.go
package main

import (
	"fmt"
)

func swap(a *float64, b *float64) {
	var temp *float64
	*temp = *a
	*a = *b
	*b = *temp
}
func main() {
	x := 3.4
	y := 4.1
	swap(&x, &y)

	fmt.Println("X is now ", x)
	fmt.Println("Y is now ", y)
}
