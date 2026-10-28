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
	l.buf.Reset()
	for {
		line, isPrefix, err := l.reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				l.err = err
			}
			return false
		}

		if _, err := l.buf.Write(line); err != nil {
			l.err = err
			return false
		}

		if !isPrefix {
			return true
		}
	}
}

func (l Lineiter) Next() string {
	return l.buf.String()
}

func (l Lineiter) Err() error {
	return l.err
}
