package main

import "fmt"

func main() {
	slice := []int {45, 23, 56, 89, 900, 26, 78}

	multiply(slice...)
}
	func multiply(b ...int) {
		ganjil := []int{}
		genap := []int{}

		for i := 0; i < len(b); i++{
			if b[i]%2 != 0 {
				ganjil = append(ganjil, b[i])
			}else{
				genap = append(genap, b[i])
			}
		}
		fmt.Println("Ganjil: ", ganjil)
		fmt.Println("Genap: ", genap)
	}