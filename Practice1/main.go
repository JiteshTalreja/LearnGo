package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := "some val"

	a, b := "again some value", 56

	fmt.Println(z, a, b)

	var c float32 = 2.3
	fmt.Println(c)

	e, f, _ := 1, 2, 3467

	fmt.Println(e, f)

}
