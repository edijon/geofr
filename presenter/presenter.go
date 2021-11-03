package presenter

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Output interface {
	Write(row []string)
}

type StandardOutput struct {
	ColumnSize uint
}

const defaultColumnSize uint = 20

func (output *StandardOutput) Write(row []string) {
	// Writes values of length output.ColumnSize (each).
	if output.ColumnSize == 0 {
		output.ColumnSize = defaultColumnSize
	}
	for i := range row {
		fmt.Printf(
			"%s",
			valueFilledWithBlanks(row[i], output.ColumnSize))
	}
	fmt.Println("") // End of line.
}

func valueFilledWithBlanks(value string, columnSize uint) string {
	// Get truncated value filled with blanks until columnSize.
	length := uint(countRunes(value))
	return truncateRunes(
		value+blankSpaces(length, columnSize),
		columnSize)
}

func countRunes(runes string) int {
	// Count runes in string, example : 'ééé' => 3
	return utf8.RuneCountInString(runes)
}

func truncateRunes(runes string, limit uint) string {
	// Equivalent of left in other languages, limit in runes.
	var i uint = 0
	var truncatedRunes string = ""
	for _, runes := range runes {
		if i >= limit {
			break
		}
		truncatedRunes += string(runes)
		i++
	}
	return truncatedRunes
}

func blankSpaces(length uint, columnSize uint) string {
	// Make blank spaces based on value length and columnSize.
	if length >= columnSize {
		return "" // Do not need any blank space.
	}
	var emptyLength uint = columnSize - length
	return strings.Repeat(" ", int(emptyLength))
}
