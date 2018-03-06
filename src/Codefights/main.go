package main

import (
	"Codefights/palindrome"
	"fmt"
)

func main() {
	if palindrome.CheckPalindrome("aaAbAaa") {
		fmt.Println("OK")
	} else {
		fmt.Println("NO")
	}

}
