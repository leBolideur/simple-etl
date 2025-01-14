package input

import (
	"encoding/csv"
	"io"
	"strconv"
)

type CellType string

const (
	CellString  CellType = "string"
	CellInt              = "int"
	CellBoolean          = "bool"
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

func CreateTableFromCSV(reader io.Reader) (*Table, error) {
	table := &Table{
		IsCompleted: false,
	}

	csvReader := csv.NewReader(reader)
	rowIndex := 0
	for {
		line, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return table, err
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

		rowIndex++
	}

	table.IsCompleted = true
	return table, nil
}

func inferCellType(rawValue string) (CellType, any) {
	intValue, err := strconv.ParseInt(rawValue, 0, 64)
	if err == nil {
		return CellInt, intValue
	}

	boolValue, err := strconv.ParseBool(rawValue)
	if err == nil {
		return CellBoolean, boolValue
	}

	return CellString, rawValue
}
