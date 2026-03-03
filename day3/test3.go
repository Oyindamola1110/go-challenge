package main

import (
	"fmt"
	"strings"
)

func main () {
	s := "hello"
	v, c := countVowelsConsonants(s)
 ratio:= vowelRatio("hello")
		  fmt.Printf("vowelRatio: %.2f\n", ratio )
	fmt.Println("Vowels:", v)
	fmt.Println("Consonants:", c)
	
}


func countVowelsConsonants (s string) (vowels int, consonants int) {

	s = strings.ToLower(s)
	for _, ch := range s {
		if  ! ((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')) {
                      continue
		}
         if ch == 'a' || ch =='e' || ch== 'i' || ch =='o' || ch == 'u' ||
		 ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			vowels++
		 }else {
			consonants++
		 }

	}
	return
}

func  vowelRatio(s string) float64 {
   vowels, consonants :=   countVowelsConsonants(s)
   total := vowels + consonants
   if total == 0 {
	return 0
   }
   return float64(vowels)/ float64(total)
}




