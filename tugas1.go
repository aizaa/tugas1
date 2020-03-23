package main

import "fmt"

func main (){
	nilai_a, nilai_b, nilai_c, nilai_d := calc(12, 34, 2345)

	fmt.Println("multiply: ", nilai_a)
	fmt.Println("divide: ", nilai_b)
	fmt.Println("add: ", nilai_c)
	fmt.Println("minus: ", nilai_d)
}
	func calc(a int, b int, c int) (int, int, int, int) {
		multiply := a * 12
		divide := multiply / 2
		add := b + 12
		minus := c - 109
		
		return multiply, divide, add, minus

}	