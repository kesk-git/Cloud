package main

import (
	"fmt"
	"io"
	"os"
)

func run(out io.Writer, args []string) error {
	if len(args) != 1 || args[0] != "--smoke" {
		return fmt.Errorf("usage: cloud-stub --smoke")
	}
	_, err := fmt.Fprintln(out, "cloud-stub: ok")
	return err
}

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
