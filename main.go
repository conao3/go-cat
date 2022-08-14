package main

import (
	"fmt"
	"io"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
)

func NewReader(filepath string) (io.ReadCloser, error) {
	var r io.ReadCloser
	switch filepath {
	case "", "-":
		r = io.NopCloser(os.Stdin)
	default:
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
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

	r, err := NewReader(opts.Input)
	if err != nil {
		return 1, err
	}
	defer r.Close()

	l := NewLineiter(r)
	for l.HasNext() {
		fmt.Println(l.Next())
	}
	if err := l.Err(); err != nil {
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
