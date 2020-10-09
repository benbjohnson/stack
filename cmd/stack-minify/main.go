package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/benbjohnson/stack"
)

func main() {
	if err := run(os.Args[1:]); err == flag.ErrHelp {
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("stack-minify", flag.ContinueOnError)
	fs.Usage = usage
	if err := fs.Parse(args); err != nil {
		return err
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	fmt.Println(string(stack.Minify(input)))
	return nil
}

func usage() {
	fmt.Println(`
stack-minify is a tool for condensing a Go stack trace read from STDIN.

Usage:

	cat mystack.txt | stack-minify

`[1:])
}
