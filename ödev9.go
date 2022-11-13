package main

import (
	"fmt"
)

func main() {
	var ArBoyut, i, KucukSira, BuyukSira int

	fmt.Print("eleman sayısını belirleyin ")
	fmt.Scan(&ArBoyut)

	Array := make([]int, ArBoyut)

	fmt.Print("\nelemanları giriniz = ")
	for i = 0; i < ArBoyut; i++ {
		fmt.Scan(&ArBoyut) //i?
	}
	EnbuyukSayi := Array[0]
	EnkucukSayi := Array[0]

	for i = 0; i < ArBoyut; i++ {
		if EnbuyukSayi < Array[i] {
			EnbuyukSayi = Array[i]
			BuyukSira = i
		}
		if EnkucukSayi > Array[i] {
			EnkucukSayi = Array[i]
			KucukSira = i
		}
	}
	fmt.Println("\nEn büyük değer    = ", EnbuyukSayi)
	fmt.Println("\nEn büyük Değerin Sıralaması = ", BuyukSira)
	fmt.Println("\nEn küçük değer= ", EnkucukSayi)
	fmt.Println("\nEn küçük değerin sıralaması = ", KucukSira)
}
