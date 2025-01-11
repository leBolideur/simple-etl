package utils

import (
	"fmt"

	"github.com/leBolideur/simple-etl/input"
)

func FindColumnIndex(colName string, header *input.Row) (int, error) {
	colIndex := -1
	for i, cell := range header.Cells {
		if cell.RawValue == colName {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		error := fmt.Errorf("no existing column with name: %s\n", colName)
		return colIndex, error
	}

	return colIndex, nil
}
