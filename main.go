package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rfratto/alphabet-echo/aecho"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var opts aecho.Options

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.BoolVar(&opts.UseWhite, "white", false, "print the not-as-good white letters instead of yellow")
	fs.BoolVar(&opts.Emphasize, "emphasize", false, "add 👏 in between words")

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	// Prefer stdin if it has data to read
	var input string
	if fi, err := os.Stdin.Stat(); err == nil && fi.Mode()&os.ModeCharDevice == 0 {
		bb, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		input = string(bb)
	} else {
		input = strings.Join(fs.Args(), ` `)
	}

	output := aecho.Transform(input, opts)
	fmt.Printf("%s\n", output)
	return nil
}
