package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println("gia tri cua a luc dau: ", a)
	// b:= a
	p := &a
	*p = 50
	fmt.Println("gia tri cua a luc sau: ",a)

	fmt.Println(&a, p)
}
