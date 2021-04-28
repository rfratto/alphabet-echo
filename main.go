package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// fragmentRegex can be used to parse out slack emoji from a word
var fragmentRegex = regexp.MustCompile("(:[a-z_+']+:|.)")

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

	for _, f := range fragmentRegex.FindAllString(normed, -1) {
		// We need to convert f to a rune slice for detecting single letter unicode
		// fragments.
		rr := []rune(f)

		switch {
		case f == ".":
			if useWhite {
				fmt.Printf(":white_circle:")
			} else {
				// The "yellow" alphabet colors are closer to orange.
				fmt.Printf(":large_orange_circle:")
			}
		case f == "?":
			fmt.Printf(":alphabet-%s-question:", selectedColor)
		case f == "@":
			fmt.Printf(":alphabet-%s-at:", selectedColor)
		case f == "#":
			fmt.Printf(":alphabet-%s-hash:", selectedColor)
		case f == "!":
			fmt.Printf(":alphabet-%s-exclamation:", selectedColor)
		case len(rr) == 1 && unicode.IsLetter(rr[0]) && rr[0] < unicode.MaxASCII:
			fmt.Printf(":alphabet-%s-%s:", selectedColor, f)
		case len(rr) == 1 && unicode.IsSpace(rr[0]):
			// Print with extra padding so words are more readable in Slack
			fmt.Printf("   ")
		default:
			// Everything else should be printed normally (this includes emoji as
			// well as unsuported individual characters)
			fmt.Printf("%s", string(f))
		}
	}
	fmt.Println()
}
