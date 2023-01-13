package main

// klavyeden girilen 2 sayının toplamını bulan program
import "fmt"

func main() {
	fmt.Print("İlk Sayıyı Girin : ")
	var sayi1 int
	fmt.Scanln(&sayi1)
	fmt.Print("İkinci Sayıyı Girin : ")
	var sayi2 int
	fmt.Scanln(&sayi2)
	var toplam = sayi1 + sayi2
	fmt.Print(toplam)
}
