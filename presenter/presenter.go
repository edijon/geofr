package presenter

import (
	"fmt"
	"strings"
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
		value := []rune(row[i]) //Helps for unicode.
		length := uint(len(value))
		fmt.Printf(
			"%s%s",
			string(value),
			string(blankSpaces(length, output.ColumnSize)))
	}
	fmt.Println("") // End of line.
}

func blankSpaces(length uint, columnSize uint) []rune {
	// Make blank spaces based on value length and columnSize.
	if length >= columnSize {
		return []rune("") // Do not need any blank space.
	}
	var emptyLength uint = columnSize - length
	return []rune(strings.Repeat(" ", int(emptyLength)))
}
