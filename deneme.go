package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const CikolataFiyata = 10
const CipsFiyata = 5
const KolaFiyata = 3

func main() {

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
		DosyadakiStok := dosyaOku("stok.txt", girilenUrun)
		if DosyadakiStok == "Stok Yok." {
			println("Bu Ürün Stoklarımızda Mevcud Değil.")
			continue
		}
		DosyadakiStokInt, _ := strconv.Atoi(DosyadakiStok)
		if DosyadakiStokInt < girilenAdet {
			fmt.Printf(" en fazla %d adet %s alabilirsiniz\n", DosyadakiStokInt, girilenUrun)
			continue
		}

		println("\nbaşka bir ürün istermisiniz")
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
				toplam += CikolataFiyata * Adet
			} else {
				toplam += CikolataFiyata*Adet - ((CikolataFiyata*Adet)*10)/100
			}

		}

		if hangiUrun == "kola" {
			toplam += KolaFiyata * Adet
		}

		if hangiUrun == "cips" {
			toplam += CipsFiyata * Adet
		}

	}

	fmt.Printf(" toplam ödenecek miktar %d TL", toplam)
}

func dosyaOku(dosyaIsmi, girilenurun string) string {
	file, err := os.Open(dosyaIsmi)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dosyadakiSatir := scanner.Text()
		arr := strings.Split(dosyadakiSatir, ":")
		if girilenurun == arr[0] {
			return arr[1]
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return "Stok Yok."
}

// sorgulama kısmı (1)

/* println("lütfen kaç adet  istediğini girin")
fmt.Scan(&girilenAdet)
if stok[girilenUrun] < (urunler[girilenUrun] + girilenAdet) {
fmt.Printf(" en fazla %d adet %s alabilirsiniz", stok[girilenUrun], girilenUrun)
continue
}  */

// stoklar

/* stok := make(map[string]int)
stok["cikolata"] = 10
stok["kola"] = 5
stok["cips"] = 3 */

// stokları stringten int e getir koşulu ver (1)

/* hesaptakiPara := dosyaOku(fmt.Sprintf("%s_hesap.txt", ad))
hesaptakiParaInt, _ := strconv.Atoi(hesaptakiPara)
toplamPara := hesaptakiParaInt + para
fmt.Printf("toplam Para %d ₺", toplamPara) */
