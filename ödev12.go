package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Println("1.sayıyı girin")
	fmt.Scanf("%d", &a)

	fmt.Println("2.sayıyı girin")
	fmt.Scanf("%d", &b)

	fmt.Println("3.sayıyı girin")
	fmt.Scanf("%d", &c)

	if (a > b) && (a > c) {
		fmt.Printf("%d,enbüyük sayi")
	}

	if (b > a) && (b > c) {
		fmt.Printf("%d,enbüyük sayi")

	} else {
		fmt.Printf("%d,enbüyük sayi")

	}

}
