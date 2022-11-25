package counter

import (
	"bytes"
	"testing"
)

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()

	args := []string{"testdata/four_lines.txt"}
	fakeTerminal := &bytes.Buffer{}
	c, err := NewCounter(
		WithInputFromArgs(args),
		WithOutput(fakeTerminal),
	)
	c.Lines()
	if err != nil {
		t.Fatal(err)
	}
	want := "4 lines\n"
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("wanted %s lines, got %s", want, got)
	}

}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("foo\nbar\nbaz\nba")
	fakeTerminal := &bytes.Buffer{}
	c, err := NewCounter(
		WithInput(inputBuf),
		WithOutput(fakeTerminal),
		WithInputFromArgs([]string{}),
	)
	c.Lines()
	if err != nil {
		t.Fatal(err)
	}
	want := "4 lines\n"
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("wanted %s lines, got %s", want, got)
	}
}
