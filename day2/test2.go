package main

import (
	"unicode"
	"fmt"
)

func main() {
	s := "Have A Good Day, 123, u0301"	
	var upper, lower, digits, spaces, other int
	for _, ch := range s {
		
		if unicode.IsUpper(ch) {
			upper ++

		} else if unicode.IsLower(ch) {
		    lower ++
		} else if unicode.IsDigit(ch) {
		    digits ++
		} else if unicode.IsSpace(ch) {
		    spaces ++
		} else {
			other ++			
		}
	}
	fmt.Println("Uppercase:" , upper)
	fmt.Println("Lowercase:" , lower)
	fmt.Println("Digits:" , digits)
	fmt.Println("Spaces:" , spaces)
    fmt.Println("Other:" , other)

}
