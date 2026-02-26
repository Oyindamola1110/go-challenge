package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input a sentence:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	runes := []rune(scanner.Text())

	if len(runes) > 0 {
		fmt.Println("First character:", string(runes[0]))
		fmt.Println("Last character:", string(runes[len(runes)-1]))
	} else {
		fmt.Println("No characters entered!")
	}

	fmt.Println("Length:", len(runes))
}
