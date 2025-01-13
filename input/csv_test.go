package input

import (
	"testing"
)

func createTestTable() *Table {
	return &Table{
		Header: &Row{
			Cells: []*Cell{
				{RawValue: "first_name"},
				{RawValue: "last_name"},
				{RawValue: "age"},
			},
			IsFiltered: false,
		},
		Rows: []*Row{
			{
				Cells: []*Cell{
					{RawValue: "John", InferedValue: "John"},
					{RawValue: "Doe", InferedValue: "Doe"},
					{RawValue: "30", InferedValue: int64(30)},
				},
				IsFiltered: false,
			},
			{
				Cells: []*Cell{
					{RawValue: "Jane", InferedValue: "Jane"},
					{RawValue: "Smith", InferedValue: "Smith"},
					{RawValue: "45", InferedValue: int64(45)},
				},
				IsFiltered: false,
			},
		},
	}
}

func TestCreateTableFromCSV(t *testing.T) {
	table, err := CreateTableFromCSV("../samples/test.csv")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTable := createTestTable()
	if len(table.Header.Cells) != len(expectedTable.Header.Cells) {
		t.Fatalf("expected %d cells in header, got=%d", len(expectedTable.Header.Cells), len(table.Header.Cells))
	}

	expectedColumnName := []string{"first_name", "last_name", "age"}
	for i, expected := range expectedColumnName {
		if table.Header.Cells[i].RawValue != expected {
			t.Fatalf("expected %q columnName, got=%q", expected, table.Header.Cells[i].RawValue)
		}
	}

	if len(table.Rows) != len(expectedTable.Rows) {
		t.Fatalf("expected %d rows, got=%d", len(expectedTable.Rows), len(table.Rows))
	}
}

func TestInferCellType_Int(t *testing.T) {
	type_, v := inferCellType("64")
	if type_ != CellInt {
		t.Fatalf("expected 'CellInt', got=%q", type_)
	}

	tt, ok := v.(int64)
	if !ok {
		t.Fatalf("expected int64, got=%T", tt)
	}

	if tt != 64 {
		t.Fatalf("expected 64, got=%d", v)
	}
}
