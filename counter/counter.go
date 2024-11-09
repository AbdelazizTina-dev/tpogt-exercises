package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	files  []io.Reader
	Input  io.Reader
	Output io.Writer
}

type Option func(*Counter) error

func WithInput(input io.Reader) Option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.Input = input
		return nil
	}
}

func WithOutput(output io.Writer) Option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output reader")
		}
		c.Output = output
		return nil
	}
}

func WithInputArgs(args []string) Option {
	return func(c *Counter) error {
		if len(args) < 1 {
			return nil
		}
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.Input = io.MultiReader(c.files...)
		return nil
	}
}

func NewCounter(opts ...Option) (*Counter, error) {
	c := &Counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func Main() int {
	c, err := NewCounter(WithInputArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0

}

func (c *Counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return lines
}
