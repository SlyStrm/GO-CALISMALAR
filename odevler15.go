package main

// Klavyeden Girilen Değerin "0" a kadar kenidisiyle toplamının sonucunu ekrana yazdıran program

import (
	"fmt"
)

var a int

var sayi int

func ToplamS(int) int {
	if sayi < 0 {
		fmt.Print(" Lütfen negatif bir değer girmeyin")
	} else {
		for i := 1; i <= sayi; i++ {
			a += (i)
		}
	}
	return a
}

func main() {
	fmt.Print("bir sayı girin: ")
	fmt.Scan(&sayi)
	fmt.Print("Toplam sonuç : ", ToplamS(sayi))
}
