package filter

import (
	"testing"

	"github.com/leBolideur/simple-etl/input"
)

func createTestTable() *input.Table {
	return &input.Table{
		Header: &input.Row{
			Cells: []*input.Cell{
				{RawValue: "first_name"},
				{RawValue: "last_name"},
				{RawValue: "age"},
			},
			IsFiltered: false,
		},
		Rows: []*input.Row{
			{
				Cells: []*input.Cell{
					{RawValue: "John", InferedValue: "John"},
					{RawValue: "Doe", InferedValue: "Doe"},
					{RawValue: "30", InferedValue: int64(30)},
				},
				IsFiltered: false,
			},
			{
				Cells: []*input.Cell{
					{RawValue: "Jane", InferedValue: "Jane"},
					{RawValue: "Smith", InferedValue: "Smith"},
					{RawValue: "45", InferedValue: int64(45)},
				},
				IsFiltered: false,
			},
		},
	}
}

func TestApplyFilter(t *testing.T) {
	table := createTestTable()

	err := ApplyFilter(table, "age:eq:30")
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if table.Rows[0].IsFiltered != false {
		t.Fatalf("expected row[0].IsFiltered: %t, got=%t", false, table.Rows[0].IsFiltered)
	}
	if table.Rows[1].IsFiltered != true {
		t.Fatalf("expected row[1].IsFiltered: %t, got=%t", false, table.Rows[1].IsFiltered)
	}
}

func TestParseIntFilter(t *testing.T) {
	input := "age:>:50,year:<:2020,age:=:50"

	filters, err := parseIntFilter(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(filters) != 3 {
		t.Fatalf("expected 3 filters, got=%d", len(filters))
	}

	expected := []struct {
		columnName  string
		filterValue int64
	}{
		{"age", 50},
		{"year", 2020},
		{"age", 50},
	}

	for i, filter := range filters {
		if filter.columnName != expected[i].columnName {
			t.Fatalf("expected columnName %q, got=%q", expected[i].columnName, filter.columnName)
		}

		if filter.filterValue != expected[i].filterValue {
			t.Fatalf("expected filterValue %d, got=%d", expected[i].filterValue, filter.filterValue)
		}
	}
}
