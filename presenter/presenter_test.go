package presenter

import (
	"testing"
)

type messageStringOutput struct {
	message string
}

func (output messageStringOutput) Write(row []string) {
	output.message = row[0]
}

func TestGivenMessageStringOutputThenIsImplementationOfOutput(t *testing.T) {
	output := messageStringOutput{}
	var i interface{} = output
	_, ok := i.(Output)
	if !ok {
		t.Errorf("An output should implements the Output interface")
	}
}

func TestGivenMessageStringOutputWhenWritesRun(t *testing.T) {
	output := messageStringOutput{}
	output.Write([]string{"test"})
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
