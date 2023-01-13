package main

//Klavyeden kenar uzunlukları girilen üçgenin ne çeşit üçgen olduğunu ve çevresini erkana yazdıran program
//

import "fmt"

func main() {

	var Kenar1 int
	var Kenar2 int
	var Kenar3 int

	fmt.Println("1.kenarı giriniz")
	fmt.Scanf("%d", &Kenar1)

	fmt.Println("2.kenarı giriniz")
	fmt.Scanf("%d", &Kenar2)

	fmt.Println("2.kenarı giriniz")
	fmt.Scanf("%d", &Kenar3)

	ucgenCevresi := cevreHesap(Kenar1, Kenar2, Kenar3)
	{
		fmt.Printf("\nüçgenin çevresi: %d", ucgenCevresi)
	}

	if (Kenar1 == Kenar2) && (Kenar2 == Kenar3) {
		fmt.Printf("Esitkenar ucgen\n")
	} else if (Kenar1 == Kenar2) || (Kenar1 == Kenar3) || (Kenar2 == Kenar3) {
		fmt.Printf("Ikizkenar ucgen\n")
	} else {
		fmt.Printf("Cesitkenar ucgen\n")
	}

	return
}

func cevreHesap(kenar1, kenar2, kenar3 int) int {
	return kenar1 + kenar2 + kenar3
}
