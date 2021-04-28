package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rfratto/alphabet-echo/aecho"
)

func main() {
	var opts aecho.Options

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.BoolVar(&opts.UseWhite, "white", false, "print the not-as-good white letters instead of yellow")
	fs.BoolVar(&opts.Emphasize, "emphasize", false, "add 👏 in between words")

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	output := aecho.Transform(strings.Join(fs.Args(), ` `), opts)
	fmt.Printf("%s\n", output)
}
