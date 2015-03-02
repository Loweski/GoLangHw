// golanghw project main.go
package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 100 {
		if i%3 == 0 {
			fmt.Println(i)
		}
		i = i + 1
	}
}
