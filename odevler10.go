package main

// Klavyeden girilecek cinsiyet, yaş ve boy değerlerine göre askere gidip gidilemeyeceğini belirleyen program

import "fmt"

const AskerlikYasi = 21

func main() {

	var Cinsiyet string
	var Yas int
	var Boy int

	fmt.Printf("Cinsiyetinizi girin: ")
	fmt.Scanf("%s", &Cinsiyet)

	fmt.Printf("Yaşınızı girin: ")
	fmt.Scanf("%d", &Yas)

	fmt.Printf("Boyunuz girin: ")
	fmt.Scanf("%d", &Boy)

	if (Cinsiyet == "Erkek" || Cinsiyet == "erkek" || Cinsiyet == "ERKEK") && (Yas >= 19 && Yas <= 40) && (Boy >= 170 && Boy <= 190) {
		fmt.Println("\nAskere gidebilirsin.")
	} else {

		if (Yas >= AskerlikYasi && Yas <= 40) == false {
			fmt.Println("\nAskere gidemezsiniz, çünkü yaşınız uygun değil.")
		}

		if (Cinsiyet == "Erkek" || Cinsiyet == "erkek" || Cinsiyet == "ERKEK") == false {
			fmt.Println("\nAskere gidemezsiniz, çünkü cinsiyetiniz uygun değil.")
		}

		if (Boy >= 170 && Boy <= 190) == false {
			fmt.Println("\nAskere gidemezsiniz, çünkü boyunuz uygun değil.")

		}
	}

}
