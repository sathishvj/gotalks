package main

import "fmt"

//show A type OMIT
type A struct {
	name string
}

//end show A type OMIT

//start main OMIT
func main() {
	fmt.Println("Hello World")
}

//end main OMIT

func printStr(s string) { // HL12
	fmt.Println(s) // HL
}
