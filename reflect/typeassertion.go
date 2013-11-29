package main

import (
	"fmt"
	"reflect"
)

//START TYPE OMIT
type A struct {
	S string // HL
}

//END TYPE OMIT

//beginning of main OMIT
func main() {
	var a A // HL123
	var in interface{}
	//2 lines into main
	in = a
	fmt.Println(in)
	var i = 5
	//5 lines into main
	in = i
	fmt.Println(in)
	v := reflect.ValueOf(in)
	fmt.Println(v)
	anotheri := in.(int)

	fmt.Println(anotheri)

}
