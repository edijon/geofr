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
		value := row[i]
		length := uint(utf8.RuneCountInString(value))
		value = truncateRunes(
			value+blankSpaces(length, output.ColumnSize),
			output.ColumnSize)
		fmt.Printf("%s", value)
	}
	fmt.Println("") // End of line.
}

func blankSpaces(length uint, columnSize uint) string {
	// Make blank spaces based on value length and columnSize.
	if length >= columnSize {
		return "" // Do not need any blank space.
	}
	var emptyLength uint = columnSize - length
	return strings.Repeat(" ", int(emptyLength))
}

func truncateRunes(value string, limit uint) string {
	var i uint = 0
	var result string = ""
	for _, value := range value {
		if i >= limit {
			break
		}
		result += string(value)
		i++
	}
	return result
}
