package main

//GÜN SAYISI  - DOĞUM
import (
	"fmt"
	"time"
)

func main() {

	var yil int
	var ay int
	var gun int
	println("Doğum yılınızı girin")
	fmt.Scan(&yil)
	println("Doğum ayınızı girin")
	fmt.Scan(&ay)
	println("Doğum tarihi gününüzü girin")
	fmt.Scan(&gun)

	dogumTarihi := DateFormat2(yil, ay, gun)
	bugun := time.Now()
	gunSayisi := bugun.Sub(dogumTarihi).Hours() / 24

	fmt.Println("Şu Kadar Gündür Hayattasınız ", gunSayisi)
}

func DateFormat2(yil, ay, gun int) time.Time {
	return time.Date(yil, time.Month(ay), gun, 0, 0, 0, 0, time.UTC)
}
