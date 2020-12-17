package main

import (
	"flag"
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
	var (
		useWhite bool
	)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.BoolVar(&useWhite, "white", false, "print the not-as-good white letters instead of yellow")

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	selectedColor := "yellow"
	if useWhite {
		selectedColor = "white"
	}

	// Normalize unicode characters to ASCII
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normed, _, _ := transform.String(t, strings.Join(fs.Args(), ` `))

	for _, c := range normed {
		switch {
		case c == '?':
			fmt.Printf(":alphabet-%s-question:", selectedColor)
		case c == '@':
			fmt.Printf(":alphabet-%s-at:", selectedColor)
		case c == '#':
			fmt.Printf(":alphabet-%s-hash:", selectedColor)
		case c == '!':
			fmt.Printf(":alphabet-%s-exclamation:", selectedColor)
		case unicode.IsLetter(c) && c < unicode.MaxASCII:
			fmt.Printf(":alphabet-%s-%s:", selectedColor, string(c))
		case unicode.IsSpace(c):
			// Print with extra padding so words are more readable in Slack
			fmt.Printf("   ")
		default:
			// Unsupported characters should just be printed normally
			fmt.Printf("%s", string(c))
		}
	}
	fmt.Println()
}
