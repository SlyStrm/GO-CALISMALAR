package main

import "fmt"

func main() {
	var n int
	var toplam int
	var i int

	fmt.Printf("1. Sayıyı Girin : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("2. Sayıyı Girin : ")
	fmt.Scanf("%d", &i)

	for i := 1; i <= n; i++ {
		toplam += i
	}

	fmt.Printf("1'den %d 'e kadar olan sayıların toplamı : %d", n, toplam)
}
