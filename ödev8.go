package main

import (
	"fmt"
)

func main() {

	var sayilar []int
	for i := 1; i <= 10; i++ {
		var sayi int
		fmt.Println("\nBir sayı giriniz...")
		fmt.Scan(&sayi)

		sayilar = append(sayilar, sayi)
	}

	enBuyukSayi := 0

	for _, girilenSayi := range sayilar {

		if girilenSayi > enBuyukSayi {
			enBuyukSayi = girilenSayi
		}
		fmt.Println(enBuyukSayi)
	}
}
