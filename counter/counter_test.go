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
	c.Input = buf
	want := 5
	got := c.Lines()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
