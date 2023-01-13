package main

// Klavyeden Girilen Değerler arasında En büyük sayı ve sırasını En küçük sayı ve sırasını Ekrana yazdıran program

import (
	"fmt"
)

func main() {
	var arBoyut, i, kucukSira, buyukSira int

	fmt.Print("eleman sayısını belirleyin ")
	fmt.Scan(&arBoyut)

	Array := make([]int, arBoyut)

	fmt.Print("\nelemanları giriniz = ")
	for i = 0; i < arBoyut; i++ {
		fmt.Scan(&arBoyut)

	}
	enBuyukSayi := Array[0]
	enKucukSayi := Array[0]

	for i = 0; i < arBoyut; i++ {
		if enBuyukSayi < Array[i] {
			enBuyukSayi = Array[i]
			buyukSira = i
		}
		if enKucukSayi > Array[i] {
			enKucukSayi = Array[i]
			kucukSira = i
		}
	}
	fmt.Println("\nEn büyük değer    = ", enBuyukSayi)
	fmt.Println("\nEn büyük Değerin Sıralaması = ", buyukSira)
	fmt.Println("\nEn küçük değer= ", enKucukSayi)
	fmt.Println("\nEn küçük değerin sıralaması = ", kucukSira)
}
