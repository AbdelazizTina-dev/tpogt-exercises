package counter

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() *Counter {
	return &Counter{
		Input: os.Stdin,
	}
}

func Main() {
	print(NewCounter().Count())
}

func (c *Counter) Count() int {
	lines := 0
	input := bufio.NewScanner(c.Input)
	for input.Scan() {
		lines++
	}
	return lines
}
