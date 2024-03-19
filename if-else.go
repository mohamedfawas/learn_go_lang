package main

import "fmt"

func main() {
	if num := 100; num < 0 {
		fmt.Println(num, "num is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}
