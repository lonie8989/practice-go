package main

import "fmt"

func main() {
	firstName := "namgi"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)
}
