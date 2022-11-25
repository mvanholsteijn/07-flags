package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type option func(*counter) error
type counter struct {
	input  io.ReadCloser
	output io.Writer
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return c, err
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		rc, ok := input.(io.ReadCloser)
		if !ok {
			rc = io.NopCloser(input)
		}
		c.input = rc
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) == 0 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func (c counter) Lines() {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	c.input.Close()
	fmt.Fprintf(c.output, "%d lines\n", lines)
}

func Lines() {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)

	// If the only person who will ever call Lines can’t do anything
	// useful with the error except log it and crash, we may as well do that directly:
	//
	// There is no useful information we can give the user,
	// because the user can’t fix our pro‐ gram, and this is
	// definitely an internal program bug. That’s exactly what panic
	// is for: reporting unrecoverable internal program bugs.
	if err != nil {
		panic(err)
	}
	c.Lines()
}
