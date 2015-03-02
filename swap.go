// golanghw project main.go
package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	var temp *int
	*temp = *a
	*a = *b
	*b = *temp
}
func main() {
	x := 3
	y := 4
	swap(&x, &y)

	fmt.Println("X is now ", x)
	fmt.Println("Y is now ", y)
}
