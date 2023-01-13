package main

// Klavyeden girilen 3 değer arasında girilen en küçük sayıyı bulan program (For Döngüsü ile)

import (
	"fmt"
)

func main() {

	var Sayilar []int
	for i := 1; i <= 3; i++ {
		var Sayi int
		fmt.Println("Bir Sayı Giriniz...")
		fmt.Scan(&Sayi)

		Sayilar = append(Sayilar, Sayi)

		if Sayi > 1000001 {
			break
		}
	}

	EnKucukSayi := 1000000

	for _, GirilenSayi := range Sayilar {

		if GirilenSayi < EnKucukSayi {
			EnKucukSayi = GirilenSayi
		}

	}
	fmt.Println(EnKucukSayi)
}
