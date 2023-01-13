package main

// SIRALAMA BÜYÜKTEN KÜÇÜĞE

import (
	"fmt"
	"sort"
)

func main() {

	var girilensayilar []int

	for i := 1; i <= 5; i++ {
		var girilensayi int

		fmt.Print("değerleri girin")

		fmt.Scan(&girilensayi)

		girilensayilar = append(girilensayilar, girilensayi)

	}

	sort.Sort(sort.Reverse(sort.IntSlice(girilensayilar)))

	fmt.Println("\ndeğerlerin sıralaması")
	fmt.Println("sıralama : ", girilensayilar)

}
