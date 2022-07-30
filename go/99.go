package main

import "fmt"

func main(){
	fmt.Println("-------9 9 乘法表-------")
	var (
		j,i  = 1,1

	)
	for i=1;i<=9;i++{
		for j=1;j<=i;j++{
			fmt.Printf("%d*%d=%-2d ",j,i,i*j)
		}
		fmt.Print("\n")
	}

	fmt.Println("======================================== 2 ===================================")
	for i=1;i<=9;i++{
		for j=i;j<=9;j++{
			fmt.Printf("%d*%d=%-2d ",i,j,i*j)
		}
		fmt.Printf("\n")
	}
}
