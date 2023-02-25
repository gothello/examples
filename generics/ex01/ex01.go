package main

import (
	"fmt"
)

func calcInt(n, x int) int {
	return n + x
}

func calcFloat(n, x float64) float64 {
	return n + x
}

type MyType interface {
	int | float64
}

func Gen[T MyType](n, x T) T {
	return n + x
}

func main() {
	ri := calcInt(1, 2)
	rf := calcFloat(1.2, 2.2)
	gif := Gen(1.56, 2.3)

	fmt.Println(ri)
	fmt.Println(rf)
	fmt.Println(gif)

}
