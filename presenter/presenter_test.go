package presenter

import (
	"testing"
)

type messageStringOutput struct {
	message string
}

func (output *messageStringOutput) Write(row []string) {
	output.message = row[0]
}

func TestGivenSliceOfStringsWhenStringOutputWriteThenWrites(t *testing.T) {
	output := &messageStringOutput{}
	output.Write([]string{"hello"})
	expected := "hello"
	if output.message != expected {
		t.Errorf("Got %q instead of %q", output.message, expected)
	}
}

type textStringOutput struct {
	text string
}

func (output *textStringOutput) Write(row []string) {
	output.text = row[0]
}

func TestGivenSliceOfStringsWhenTextOutputWriteThenWrites(t *testing.T) {
	output := &textStringOutput{}
	output.Write([]string{"world"})
	expected := "world"
	if output.text != expected {
		t.Errorf("Got %q instead of %q", output.text, expected)
	}
}

func TestGivenSliceOfStringsWhenStandardOutputWriteThenRun(t *testing.T) {
	output := &StandardOutput{}
	output.Write([]string{"Hello", "World"})
}
