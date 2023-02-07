package main

// Klavyeden girilen 3 değer arasında girilen en büyük sayıyı bulan program

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Println("1.Sayıyı Girin")
	fmt.Scanf("%d", &a)

	fmt.Println("2.Sayıyı Girin")
	fmt.Scanf("%d", &b)

	fmt.Println("3.Sayıyı Girin")
	fmt.Scanf("%d", &c)

	if c == 2 {
		dosyayaYaz("banka.txt", "cipse kampanya uygulandı\n")
	}

	if (a > b) && (a > c) {
		fmt.Printf("%d,En Büyük Sayi", a)
	}

	if (b > a) && (b > c) {
		fmt.Printf("%d,En Büyük Sayi", b)

	} else {
		fmt.Printf("%d,En Büyük Sayi", c)

	}

}
func dosyayaYaz(dosyaIsmi string, yazi string) {
	f, err := os.OpenFile(dosyaIsmi, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(yazi); err != nil {
		panic(err)
	}
}
