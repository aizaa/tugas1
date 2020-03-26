package main

import "fmt"

//soal pertama
func anonymous_func() func (string) string {
	return func (msg string) string {
		return msg
	}
}

func main(){

	//soal kedua
	nilai := []int {45, 23, 56, 89, 900, 26, 78}
	
	var GanjilGenap = func(a ...int) ([]int, []int){
	var ganjil = []int{}
	var genap = []int{}

		for _, a := range a {
			if a%2 == 0 {
				ganjil = append(ganjil, a)
			}else{
				genap = append(genap, a)
			}
		}
		return ganjil, genap
	}
	var ganjil, genap = GanjilGenap(nilai...)
	
	fmt.Println("Ganjil: ", ganjil)
	fmt.Println("Genap: ", genap)

	print_func := anonymous_func()
	fmt.Println(print_func("Hi there anonymous"))
}