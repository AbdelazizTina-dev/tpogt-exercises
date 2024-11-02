package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
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

func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Lines())
}

func (c *Counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	return lines
}
