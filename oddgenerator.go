// golanghw project main.go
package main

import (
	"fmt"
)

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i + 1
		i += 2
		return
	}
}
func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 0
	fmt.Println(nextOdd()) // 2
	fmt.Println(nextOdd()) // 4
}
