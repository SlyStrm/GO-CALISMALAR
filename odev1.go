package main

// klavyeden girilen 3 notun ortalamasını alan program
import (
	"fmt"
)

func main() {

	var (
		not1 int
		not2 int
		not3 int
	)

	fmt.Println("Lütfen notlarınızı giriniz;")

	fmt.Scan(&not1)

	fmt.Scan(&not2)

	fmt.Scan(&not3)

	notOrtalamasi := ortalamaAl(not1, not2, not3)

	if notOrtalamasi < 45 {
		fmt.Printf("\nNot ortalamanız: %d Yani çok düşük..", notOrtalamasi)

	} else if notOrtalamasi > 70 && notOrtalamasi < 100 {
		fmt.Printf("\nNot ortalamanız: %d Yani iyi..", notOrtalamasi)

	} else if notOrtalamasi > 46 && notOrtalamasi < 69 {
		fmt.Printf("\nNot ortalamanız: %d Yani orta..", notOrtalamasi)
	}

}

func ortalamaAl(not1, not2, not3 int) int {
	return not1*10/100 + not2*30/100 + not3*60/100
}
