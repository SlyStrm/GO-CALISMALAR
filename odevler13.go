package main

// Klavyeden Girilen Değerler ile 4 İşlem Yapan Program (switch case ile)

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var operator string

	fmt.Println("işlemi yapılacak sayıyı  girin")
	fmt.Scanf("%d", &a)

	fmt.Println("işlemi yapılacak sayıyı girin")
	fmt.Scanf("%d", &b)

	fmt.Println("işlem biçimini girin")
	fmt.Scanf("%s", &operator)

	switch operator {
	case "+":
		fmt.Printf("%d %s %d = %d", a, operator, b, a+b)
	case "-":
		fmt.Printf("%d %s %d = %d", a, operator, b, a-b)
	case "*":
		fmt.Printf("%d %s %d = %d", a, operator, b, a*b)
	case "%":
		var Sonuc float64

		Sonuc = float64(a)/100*10 + float64(b)/100*10

		fmt.Printf("%d, %s ,%d = %f", a, operator, b, Sonuc)
	case "/":
		if a == 0 || b == 0 {
			fmt.Printf("\n0 a Bölüm Her Zaman 0 Çıkar Lütfen Başka Bir Sayı Girin")
		} else {

			var sonuc float64
			sonuc = float64(a) / float64(b)

			fmt.Printf("\n %d %s %d = %0.2f", a, operator, b, sonuc)
		}
	}
}
