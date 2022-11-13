package main

import "fmt"

func main() {
	var n int
	var toplam int

	fmt.Printf("N sayısını Girin : ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		toplam += i
	}

	fmt.Printf("1'den %d 'e kadar olan sayıların toplamı : %d", n, toplam)
}
