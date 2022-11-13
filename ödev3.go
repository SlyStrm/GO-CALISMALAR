package main

import "fmt"

func main() {
	// 0 - 100 arasındaki tek sayıları ekrana yazdıran programı yaz
	// sonra ürettiğin  sayıları tekrar bir for döngüsüne sok ve her bir sayıyı tekmi mi diye kontrol et ve yazdır.

	var sifirlaYuzArasindakiSayilar []int

	for s := 1; s <= 100; s++ {
		sifirlaYuzArasindakiSayilar = append(sifirlaYuzArasindakiSayilar, s)
	}

	for _, sayi := range sifirlaYuzArasindakiSayilar {

		kalan := sayi % 2

		if kalan != 0 {
			fmt.Printf("Tek sayı: %d \n", sayi)
		}
	}

}
