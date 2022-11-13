package main

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
		var sonuc float64

		sonuc = float64(a)/100*10 + float64(b)/100*10

		fmt.Printf("%d, %s ,%d = %f", a, operator, b, sonuc)
	case "/":
		if a == 0 || b == 0 {
			fmt.Printf("\n0 a bölüm her zaman 0 çıkar lütfen başka bir sayı girin")
		} else {

			var sonuc float64
			sonuc = float64(a) / float64(b)

			fmt.Printf("\n %d %s %d = %0.2f", a, operator, b, sonuc)
		}
	}
}
