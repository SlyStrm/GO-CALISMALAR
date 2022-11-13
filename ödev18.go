package main

import (
	"fmt"
)

//Sayının faktöriyelini bulan programa

var fact float64 = 1

var i int = 1
var sayi int

func factorial(n int) float64 {
	if sayi < 0 {
		fmt.Print("Factoriyel sayı negatif olmaz.")
	} else {
		for i := 1; i <= sayi; i++ {
			fact *= float64(i)
		}
	}
	return fact
}

func main() {
	fmt.Print("bir sayı girin: ")
	fmt.Scan(&sayi)
	fmt.Print("Factoriyel sonuç : ", factorial(sayi))
}
