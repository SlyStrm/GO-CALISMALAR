package main

//klavyeden girilen değerler ile 4 işlem yapan program

import "fmt"

func main() {
	var sayi1 int
	var sayi2 int
	var toplama int
	var carpma int
	var bolme int
	var cikarma int

	fmt.Println("1.Sayıyı Girin")
	fmt.Scanln(&sayi1)
	fmt.Println("2.Sayıyı Girin")
	fmt.Scanln(&sayi2)
	toplama = (sayi1) + sayi2
	carpma = sayi1 * sayi2
	cikarma = sayi1 - sayi2
	bolme = sayi1 / sayi2
	fmt.Printf("\ntoplam=%d\n", toplama)
	fmt.Printf("carpma=%d\n", carpma)
	fmt.Printf("fark=%d\n", cikarma)
	fmt.Printf("bolme= %d\n", bolme)
}
