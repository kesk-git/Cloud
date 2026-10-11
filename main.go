package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const usage = "usage: cloud-stub --smoke [--name NAME]"

// version is printed by "--version" as "cloud-stub <version>".
const version = "0.2.0"

// parseSmoke validates the command line and returns the optional smoke name.
// Accepted forms are "--smoke" and "--smoke --name NAME". NAME has its
// surrounding whitespace trimmed; an empty or whitespace-only NAME is an error.
func parseSmoke(args []string) (string, error) {
	if len(args) == 0 || args[0] != "--smoke" {
		return "", errors.New(usage)
	}
	switch len(args) {
	case 1:
		return "", nil
	case 3:
		if args[1] != "--name" {
			return "", errors.New(usage)
		}
		name := strings.TrimSpace(args[2])
		if name == "" {
			return "", errors.New("cloud-stub: --name requires a nonempty name")
		}
		return name, nil
	default:
		return "", errors.New(usage)
	}
}

func run(out io.Writer, args []string) error {
	if len(args) == 1 && args[0] == "--version" {
		_, err := fmt.Fprintln(out, "cloud-stub "+version)
		return err
	}
	name, err := parseSmoke(args)
	if err != nil {
		return err
	}
	line := "cloud-stub: ok"
	if name != "" {
		line += " (" + name + ")"
	}
	_, err = fmt.Fprintln(out, line)
	return err
}

func main() {
	if err := run(os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
