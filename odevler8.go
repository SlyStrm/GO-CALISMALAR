package main

// Klavyeden girilen 3 değer arasında girilen en büyük sayıyı bulan program

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Println("1.Sayıyı Girin")
	fmt.Scanf("%d", &a)

	fmt.Println("2.Sayıyı Girin")
	fmt.Scanf("%d", &b)

	fmt.Println("3.Sayıyı Girin")
	fmt.Scanf("%d", &c)

	if (a > b) && (a > c) {
		fmt.Printf("%d,En Büyük Sayi", a)
	}

	if (b > a) && (b > c) {
		fmt.Printf("%d,En Büyük Sayi", b)

	} else {
		fmt.Printf("%d,En Büyük Sayi", c)

	}

}
