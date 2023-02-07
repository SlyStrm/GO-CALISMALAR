package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dosyadakiKelime := dosyaOku2(fmt.Sprintf("yazı.txt"))

	inputs := []string{"alm", "Leopard"}

	for _, input := range inputs {
		if strings.Contains(dosyadakiKelime, input) {
			str := dosyadakiKelime
			yeniDeger := strings.Repeat("*", len(input))
			n := -1
			s := strings.Replace(str, input, yeniDeger, n)
			dosyadakiKelime = s
		} else {
			println("Aranan Kelime Mevcud Değil!")
		}

	}

	fmt.Println(dosyadakiKelime)
}

func dosyaOku2(dosyaIsmi string) string {
	file, err := os.Open(dosyaIsmi)
	if err != nil {
		log.Fatal(err)
	}

	text := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text() + " "
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return text
}
