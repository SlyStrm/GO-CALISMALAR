package main

import "fmt"

func main() {
	// Soru: Kullanıcıdan cinsiyet ve yaş al Eğer cinsiyeti erkekse(erkek,ERKEK,Erkek)  yaşı 19-40 arasıysa ve boyu 170-190 arasıysa askere gidebilirsin yazdır.

	var cinsiyet string
	var yas int
	var boy int

	fmt.Printf("Cinsiyetinizi girin: ")
	fmt.Scanf("%s", &cinsiyet)

	fmt.Printf("Yaşınızı girin: ")
	fmt.Scanf("%d", &yas)

	fmt.Printf("Boyunuz girin: ")
	fmt.Scanf("%d", &boy)

	if (cinsiyet == "Erkek" || cinsiyet == "erkek" || cinsiyet == "ERKEK") && (yas >= 19 && yas <= 40) && (boy >= 170 && boy <= 190) {
		fmt.Println("\nAskere gidebilirsin.")
	} else {

		if (yas >= 19 && yas <= 40) == false {
			fmt.Println("\nAskere gidemezsiniz, çünkü yaşınız uygun değil.")
		}

		if (cinsiyet == "Erkek" || cinsiyet == "erkek" || cinsiyet == "ERKEK") == false {
			fmt.Println("\nAskere gidemezsiniz, çünkü cinsiyetiniz uygun değil.")
		}

		if (boy >= 170 && boy <= 190) == false {
			fmt.Println("\nAskere gidemezsiniz, çünkü boyunuz uygun değil.")
		}
	}

}
