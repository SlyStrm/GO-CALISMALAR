package main

import (
	"fmt"
	"strconv"
)

func main() {

	var Array []int

	for s := 1; s <= 100; s++ {

		Array = append(Array, s)
	}
	for _, sayi := range Array {

		kalan := sayi % 2
		if kalan == 0 && sayi > 9 {
			slice := strconv.Itoa(sayi)
			{
				if slice[1] == 0 {
					fmt.Println(sayi)
				}

			}
		}
	}
}
