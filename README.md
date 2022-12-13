# KODLAMA
-package main

import (
	"fmt"
	"time"
)

func main() {

	var yıl int
	var ay int
	var gun int

	fmt.Scanf(&yıl)
	fmt.Scanf(&ay)
	fmt.Scanf(&gun)

	z1 := Date(yıl, ay, gun)
	z2 := Date(2022, 11, 13) // time.Now() ?
	gunler := z2.Sub(z1).Hours() / 24
	fmt.Println("şu kadar gündür hayattasınız o7", gunler)
}

func Date(year, month, day int) time.Time {
	return time.Date(yıl, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
