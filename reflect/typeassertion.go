package main

import "fmt"

//start OMIT
type A struct {
	S string
}

func main() {
	var a A // HL123
	var in interface{}
	in = a

	var aAgain A    // HL
	aAgain = in.(A) // HL
	fmt.Println(aAgain)

	arr := []int{5, 6, 7, 8}
	in = arr
	var arrAgain []int    // HL
	arrAgain = in.([]int) // HL
	fmt.Println(arrAgain)
}

//end OMIT
