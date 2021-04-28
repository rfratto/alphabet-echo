// Package aecho (pronounced "echo") allows for transforming input
// to Slack's Alphabet emoji form.
package aecho

import (
	"fmt"
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

type Options struct {
	UseWhite bool
}

// Transform will transform the input string into an Alphabetized version.
func Transform(input string, opts Options) string {
	selectedColor := "yellow"
	if opts.UseWhite {
		selectedColor = "white"
	}

	var w strings.Builder

	// Normalize unicode characters to ASCII
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normed, _, _ := transform.String(t, string(input))

	for _, f := range fragmentRegex.FindAllString(normed, -1) {
		// We need to convert f to a rune slice for detecting single letter unicode
		// fragments.
		rr := []rune(f)

		switch {
		case f == ".":
			if opts.UseWhite {
				fmt.Fprintf(&w, ":white_circle:")
			} else {
				// The "yellow" alphabet colors are closer to orange.
				fmt.Fprintf(&w, ":large_orange_circle:")
			}
		case f == "?":
			fmt.Fprintf(&w, ":alphabet-%s-question:", selectedColor)
		case f == "@":
			fmt.Fprintf(&w, ":alphabet-%s-at:", selectedColor)
		case f == "#":
			fmt.Fprintf(&w, ":alphabet-%s-hash:", selectedColor)
		case f == "!":
			fmt.Fprintf(&w, ":alphabet-%s-exclamation:", selectedColor)
		case len(rr) == 1 && unicode.IsLetter(rr[0]) && rr[0] < unicode.MaxASCII:
			fmt.Fprintf(&w, ":alphabet-%s-%s:", selectedColor, f)
		case len(rr) == 1 && unicode.IsSpace(rr[0]):
			// Print with extra padding so words are more readable in Slack
			fmt.Fprintf(&w, "   ")
		default:
			// Everything else should be printed normally (this includes emoji as
			// well as unsuported individual characters)
			fmt.Fprintf(&w, "%s", string(f))
		}
	}

	return w.String()
}
