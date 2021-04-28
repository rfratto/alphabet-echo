package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rfratto/alphabet-echo/aecho"
)

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

	output := aecho.Transform(strings.Join(fs.Args(), ` `), aecho.Options{
		UseWhite: useWhite,
	})
	fmt.Printf("%s\n", output)
}
