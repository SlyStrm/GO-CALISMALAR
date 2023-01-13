package main

// SIRALAMA KÜÇÜKTEN - BÜYÜĞE

import (
	"fmt"
	"sort"
)

func main() {

	var girilenSayilar []int

	for i := 1; i <= 5; i++ {
		var girilenSayi int
		// sort.Ints (scl2)
		fmt.Print("değerleri girin")

		fmt.Scan(&girilenSayi)

		girilenSayilar = append(girilenSayilar, girilenSayi)

	}
	sort.Ints(girilenSayilar)

	fmt.Println("\ndeğerlerin sıralaması")
	fmt.Println("sıralama : ", girilenSayilar)

}
