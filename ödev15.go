package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Println("\n1.sayıyı girin")
	fmt.Scanf("%d", &a)

	fmt.Println("\n2.sayıyı girin")
	fmt.Scanf("%d", &b)

	fmt.Println("\n3.sayıyı girin")
	fmt.Scanf("%d", &c)

	if (a < b) && (a < c) {
		fmt.Printf("en küçük sayi\n")
	}

	if (b < a) && (b < c) {
		fmt.Printf("%d en küçük sayi\n")

	} else {
		fmt.Printf("%d en küçük sayi\n")

	}

}
