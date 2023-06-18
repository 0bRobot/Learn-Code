package main

import "fmt"

func main() {

	//demo()
	demo1()

}
func demo() {
	sum := 0
	var a = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]

	}
	fmt.Println(sum)
}

func demo1() {
	var b = [...]int{1, 3, 5, 7, 8}
	ch := 8
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[i]+b[j] == ch {
				fmt.Println(i, j)
			}
		}
	}

}
