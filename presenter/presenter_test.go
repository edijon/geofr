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

func TestGivenUnicodeValueWhenCountRunesThenGetNumberOfRunes(t *testing.T) {
	unicodeValue := "Pyrén"
	expected := 5
	current := countRunes(unicodeValue)
	if current != expected {
		t.Errorf("Got %q instead of %q", current, expected)
	}
}

func TestGivenUnicodeStringAndLimit10WhenTruncateThenTruncate10Runes(t *testing.T) {
	unicodeValue := "Pyrénées-Orientales"
	expected := "Pyrénées-O"
	current := truncateRunes(unicodeValue, 10)
	if current != expected {
		t.Errorf("Got %q instead of %q", current, expected)
	}
}

func TestGivenStringWhenValueFilledWithBlanksGetValue(t *testing.T) {
	value := "test"
	var columnSize uint = 10
	expected := "test      "
	current := valueFilledWithBlanks(value, columnSize)
	if current != expected {
		t.Errorf("Got %q instead of %q", current, expected)
	}
}
