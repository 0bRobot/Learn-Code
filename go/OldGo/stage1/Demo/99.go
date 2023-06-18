package main

import (
	"fmt"
)

func main() {

	C99()

	C991()

	C992()
}

func C99() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func C991() {
	j, i := 1, 1

	for i = 1; i <= 9; i++ {
		for j = i; j <= 9; j++ {
			fmt.Printf("%d*%d=%-2d ", i, j, i*j)
		}
		fmt.Printf("\n")

	}
	fmt.Println()
}

func C992() {

	for i := 9; i >= 1; i-- {
		for j := 9; j >= i; j-- {
			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}

}
