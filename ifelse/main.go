package main

import "fmt"

var cat int = 52

func main() {

	if cat >= 80 {
		fmt.Println("cat is above 80")
	} else if cat <= 50 {
		fmt.Println("cat is below 50")
	} else {
		fmt.Printf("cat is %d\n", cat)
	}
}
