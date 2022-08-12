package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
)

func file_or_stdin(filepath string) (io.Reader, error) {
	var r io.Reader
	switch filepath {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}

	return r, nil
}

func main_() (int, error) {
	var opts struct {
		Input string `short:"i" long:"input" description:"Input file" default:"-"`
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		return 1, err
	}

	if len(args) > 0 {
		return 1, fmt.Errorf("unexpected argument: %q", args[0])
	}

	r, err := file_or_stdin(opts.Input)
	if err != nil {
		return 1, err
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return 1, err
	}

	return 0, nil
}

func main() {
	ret, err := main_()
	if err != nil {
		log.Println(err)
		os.Exit(ret)
	}
}
