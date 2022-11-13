package main

import "fmt"

func main() {

	var taban int
	var yukseklik int
	var alan float64

	fmt.Printf("\nUcgenin tabanini giriniz : ")
	fmt.Scanf("%d", &taban)

	fmt.Printf("\nUcgenin yuksekligini giriniz : ")
	fmt.Scanf("%d", &yukseklik)

	alan = float64(taban*yukseklik) / 2
	fmt.Printf("\nUcgenin alani : %.2f", alan)

}
