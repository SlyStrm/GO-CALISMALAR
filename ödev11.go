package main

import "fmt"

func main() {

	var kenar1 int
	var kenar2 int
	var kenar3 int

	fmt.Println("1.kenarı giriniz")
	fmt.Scanf("%d", &kenar1)

	fmt.Println("2.kenarı giriniz")
	fmt.Scanf("%d", &kenar2)

	fmt.Println("2.kenarı giriniz")
	fmt.Scanf("%d", &kenar3)

	if (kenar1 == kenar2) && (kenar2 == kenar3) {
		fmt.Printf("Esitkenar ucgen\n")
	} else if (kenar1 == kenar2) || (kenar1 == kenar3) || (kenar2 == kenar3) {
		fmt.Printf("Ikizkenar ucgen\n")
	} else {
		fmt.Printf("Cesitkenar ucgen\n")
	}

	return
}
