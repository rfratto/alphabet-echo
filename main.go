package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func main() {
	t := transform.Chain(norm.NFD,transform.RemoveFunc(isMn),norm.NFC)
	normed, _, _ := transform.String(t, strings.Join(os.Args[1:],` `))

	words := strings.Split(normed, ` `)
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

	fmt.Println()
}
