package main

import "fmt"

func main() {

	var girilenSayilar []int

	for i := 1; i <= 10; i++ {
		var girilenSayi int

		fmt.Print("değerleri girin")

		fmt.Scan(&girilenSayi)

		if girilenSayi > 10000000 {
			break
		}

		girilenSayilar = append(girilenSayilar, girilenSayi)

	}

	girilenToplamSayiAdedi := boyut(girilenSayilar)
	girilenEnBuyukSayi := enbuyuksayi(girilenSayilar)
	girilenEnKucukSayi := enkucuksayi(girilenSayilar)
	fmt.Printf("\ngirilen toplam sayı adeti = %d", girilenToplamSayiAdedi)
	fmt.Printf("\ngirilen en büyük sayı = %d", girilenEnBuyukSayi)
	fmt.Printf("\ngirilen en küçük sayı = %d", girilenEnKucukSayi)
}

func boyut(girilensayilar []int) int {

	return len(girilensayilar)
}

func enbuyuksayi(girilensayilar []int) int {
	enbuyuksayi := 1

	for _, girilensayi := range girilensayilar {

		if girilensayi > enbuyuksayi {
			enbuyuksayi = girilensayi
		}

	}
	return (enbuyuksayi)
}
func enkucuksayi(girilensayilar []int) int {
	enkucuksayi := 10000001
	for _, girilensayi := range girilensayilar {

		if girilensayi < enkucuksayi {
			enkucuksayi = girilensayi

		}

	}
	return (enkucuksayi)
}
