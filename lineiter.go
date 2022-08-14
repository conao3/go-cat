package main

import (
	"bufio"
	"bytes"
	"io"
)

type Lineiter struct {
	reader *bufio.Reader
	line   string
	err    error
	buf    bytes.Buffer
}

func NewLineiter(reader io.Reader) *Lineiter {
	return &Lineiter{reader: bufio.NewReader(reader)}
}

func (l *Lineiter) HasNext() bool {
	var buf bytes.Buffer
	l.buf = buf
	for {
		line, isPrefix, err := l.reader.ReadLine()

		_, err2 := buf.Write(line)
		if err2 != nil {
			l.err = err2
			return false
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			l.err = err
			return false
		}

		if !isPrefix {
			break
		}
	}

	return true
}

func (l Lineiter) Next() string {
	return l.buf.String()
}

func (l Lineiter) Err() error {
	return l.err
}
