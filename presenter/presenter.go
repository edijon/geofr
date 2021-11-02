package presenter

import (
	"fmt"
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
		value = value + blankSpaces(uint(len(value)), output.ColumnSize)
		fmt.Print(value[0:output.ColumnSize])
	}
	fmt.Println("") // End of line.
}

func blankSpaces(length uint, columnSize uint) string {
	// Make blank spaces based on value length and columnSize.
	if length >= columnSize {
		return "" // Do not need any blank space.
	}
	var emptyLength uint = columnSize - length
	var result string = ""
	var i uint = 0
	for i < emptyLength {
		result = result + " "
		i++
	}
	return result
}
