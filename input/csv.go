package input

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type CellType string

const (
	CellString CellType = "string"
	CellInt             = "int"
)

type Row struct {
	Index      int
	Cells      []*Cell
	IsFiltered bool
}

func newRow(index int, cells []*Cell, isFiltered bool) *Row {
	row := &Row{
		Index:      index,
		Cells:      cells,
		IsFiltered: isFiltered,
	}

	return row
}

type Column struct {
}

type Cell struct {
	Index        int
	RawValue     string
	Type         CellType
	InferedValue any
}

func newCell(index int, rawValue string, t CellType, inferedValue any) *Cell {
	cell := &Cell{
		Index:        index,
		RawValue:     rawValue,
		Type:         t,
		InferedValue: inferedValue,
	}

	return cell
}

type Table struct {
	Header  *Row
	Rows    []*Row
	Columns []Column

	IsCompleted bool
}

func ReadCSV(filepath string) (*Table, error) {
	table := &Table{
		IsCompleted: false,
	}

	file, err := os.Open(filepath)
	if err != nil {
		return table, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rowIndex := 0
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		cells := make([]*Cell, 0, len(line))
		for i, value := range line {
			t, v := inferCellType(value)
			cell := newCell(i, value, t, v)
			cells = append(cells, cell)
		}

		row := newRow(rowIndex, cells, false)
		if rowIndex == 0 {
			table.Header = row
		} else {
			table.Rows = append(table.Rows, row)
		}
		table.IsCompleted = true

		rowIndex++
	}

	// fmt.Println(table)
	return table, nil
}

func inferCellType(rawValue string) (CellType, any) {
	intValue, err := strconv.ParseInt(rawValue, 0, 64)
	if err == nil {
		return CellInt, intValue
	}

	return CellString, rawValue
}
