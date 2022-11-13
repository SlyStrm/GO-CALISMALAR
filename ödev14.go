package main

import (
	"fmt"
)

func main() {

	var sayilar []int
	for i := 1; i <= 3; i++ {
		var sayi int
		fmt.Println("Bir sayı giriniz...")
		fmt.Scan(&sayi)

		sayilar = append(sayilar, sayi)

		if sayi > 1000001 {
			break
		}
	}

	enKucukSayi := 1000000

	for _, girilenSayi := range sayilar {

		if girilenSayi < enKucukSayi {
			enKucukSayi = girilenSayi
		}
		fmt.Println(enKucukSayi)
	}
}
