package main

import (
	"fmt"
)

const CikolataFiyat = 10
const CipsFiyat = 5
const KolaFiyat = 3

func main() {

	stok := make(map[string]int)
	stok["cikolata"] = 10
	stok["kola"] = 5
	stok["cips"] = 3

	urunler := make(map[string]int)

	var girilenUrun string
	var girilenAdet int
	var cikis string

	for {
		println("lütfen hangi ürünü almak  istediğini girin")
		fmt.Scan(&girilenUrun)
		if girilenUrun == "ÇİKOLATA" || girilenUrun == "CİKOLATA" {
			girilenUrun = "cikolata"
		}

		if girilenUrun == "KOLA" || girilenUrun == "Kola" {
			girilenUrun = "kola"
		}

		if girilenUrun == "CİPS" || girilenUrun == "Cips" {
			girilenUrun = "cips"
		}

		if girilenUrun != "cips" && girilenUrun != "kola" && girilenUrun != "cikolata" {
			fmt.Println("yalnızca cips,kola ve cikolata alabilirsiniz")
			continue
		}

		println("lütfen kaç adet  istediğini girin")
		fmt.Scan(&girilenAdet)
		if stok[girilenUrun] < (urunler[girilenUrun] + girilenAdet) {
			fmt.Printf(" en fazla %d adet %s alabilirsiniz", stok[girilenUrun], girilenUrun)
			continue
		}

		println("başka bir ürün istermisiniz")
		fmt.Scan(&cikis)

		urunler[girilenUrun] = urunler[girilenUrun] + girilenAdet

		if cikis == "Hayır" {
			break
		}

	}

	toplam := 0

	for hangiUrun, Adet := range urunler {

		if hangiUrun == "cikolata" {

			if girilenAdet < 3 {
				toplam += CikolataFiyat * Adet
			} else {
				toplam += CikolataFiyat*Adet - ((CikolataFiyat*Adet)*10)/100
			}

		}

		if hangiUrun == "kola" {
			toplam += KolaFiyat * Adet
		}

		if hangiUrun == "cips" {
			toplam += CipsFiyat * Adet
		}

	}

	fmt.Printf(" toplam ödenecek miktar %d TL", toplam)
}
