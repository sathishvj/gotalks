package main

import (
	"fmt"
	"time"
)

func main() {
	//start OMIT
	var emptyIn interface{}
	emptyIn = 5
	emptyIn = time.Time{}
	emptyIn = []int{5, 6, 7, 8}
	//end OMIT

	fmt.Println(emptyIn)
}
