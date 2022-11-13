package main

import "fmt"

func main() {
	var n int
	var toplam int

	fmt.Printf("N sayısını Giriniz : ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= 5; i++ {
		toplam += i
	}

	fmt.Printf("1'den %d 'e kadar olan sayıların toplamı : %d", n, toplam)
}
