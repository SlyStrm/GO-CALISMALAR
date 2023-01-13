package main

// yaş gün

import "fmt"
import "time"

func main() {

	var yıl int
	var gun int

	fmt.Println(" Doğum Yılınızı ve günü girin")
	fmt.Print(" - ")
	fmt.Scanln(&yıl, &gun)

	var mevcudZaman int
	mevcudZaman = time.Now().Year()
	var yas int
	yas = mevcudZaman - yıl
	fmt.Println(yas)

	var mevcudZamanG int
	mevcudZamanG = time.Now().YearDay()
	var gunS int
	gunS = mevcudZamanG - gun
	fmt.Println(gunS)

}
