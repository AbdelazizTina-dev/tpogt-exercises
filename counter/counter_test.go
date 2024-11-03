package counter_test

import (
	"bytes"
	"testing"

	"github.com/AbdelazizTina-dev/tpogt-exercises/counter"
)

func TestCounterCountLinesFromInput(t *testing.T) {
	t.Parallel()
	buf := bytes.NewBufferString("Hi\nMy Name is\nRudy\nI am from the Greyrat Family\nNice to meet you!")
	c, err := counter.NewCounter(counter.WithInput(buf))
	if err != nil {
		t.Fatal(err)
	}
	want := 5
	got := c.Lines()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestCounterCountLinesFromInputGivenByArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := counter.NewCounter(counter.WithInputArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCounterIgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	buf := bytes.NewBufferString("Hi\nMy Name is\nRudy\nI am from the Greyrat Family\nNice to meet you!")
	c, err := counter.NewCounter(counter.WithInput(buf), counter.WithInputArgs([]string{}))
	if err != nil {
		t.Fatal(err)
	}
	want := 5
	got := c.Lines()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
