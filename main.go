package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "namgi"
	fmt.Println(*firstName)
}
