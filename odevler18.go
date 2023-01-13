package main

import "fmt"
import "time"

func main() {

	var ad string
	var soyad string
	var yıl int

	fmt.Println("Ad - SoyAdınızı - Doğum Yılınızı girin")

	fmt.Scanln(&ad, &soyad, &yıl)

	fmt.Print(soyad)
	fmt.Print("- ")

	fmt.Print(ad)
	fmt.Print("- ")

	var mevcudZaman int
	mevcudZaman = time.Now().Year()

	var yas int
	yas = mevcudZaman - yıl

	fmt.Println(yas)

}
