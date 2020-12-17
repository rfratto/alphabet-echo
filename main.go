package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for _, p := range os.Args[1:] {
		words := strings.Split(p, " ")
		for _, word := range words {
			for _, c := range word {
				letter := string(c)
				switch {
				case letter == "?":
					letter = "question"
				case letter == "@":
					letter = "at"
				case letter == "#":
					letter = "hash"
				case letter == "!":
					letter = "exclamation"
				case !unicode.IsLetter(c):
					fmt.Printf("%s", letter)
					continue
				}
				fmt.Printf(":alphabet-yellow-%s:", letter)
			}
			fmt.Printf("   ")
		}
	}
	fmt.Println()
}
