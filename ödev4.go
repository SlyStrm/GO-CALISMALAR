package main

import "fmt"

func main() {

	var girilensayilar []int

	for i := 1; i <= 10; i++ { //for i := 1; i <= n; i++ değişken "n" ise istenilecek n kadar sayı girilebilir
		var girilensayi int

		fmt.Print("değerleri girin")

		fmt.Scan(&girilensayi)

		if girilensayi > 10000000 {
			break
		}

		girilensayilar = append(girilensayilar, girilensayi)

	}

	girilenToplamSayiAdedi := boyut(girilensayilar)
	girilenEnBuyukSayi := enbuyuksayi(girilensayilar)
	girilenEnKucukSayi := enkucuksayi(girilensayilar)
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
