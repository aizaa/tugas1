package main

import "fmt"

func main() {
	mul := calc (16)

	nilai := []int {21, 31, 41}
	nilai[0] = 46
	nilai[1] = 62
	nilai[2] = 82


	fmt.Println("multiply: ", mul)
	fmt.Println(nilai)
}

	func calc(a int) (int) {
		multiply := a * 54 + 23 / 12 - 130

		return multiply
	}