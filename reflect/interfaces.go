package main

import "fmt"

//START INTERFACES OMIT
type MyInterface interface {
	AMethod()
}

type MyStruct struct {
	Id int
}

func (m MyStruct) AMethod() {}

func main() {
	var st MyStruct
	var in MyInterface
	in = st
	fmt.Println(in)
}

//END INTERFACES OMIT
