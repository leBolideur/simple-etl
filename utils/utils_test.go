package utils

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
		Rows: []*input.Row{},
	}
}

func TestFindColumnIndex(t *testing.T) {
	table := createTestTable()

	expectedIdx := map[string]int{
		"first_name": 0,
		"last_name":  1,
		"age":        2,
	}
	columns := []string{"first_name", "last_name", "age"}
	for _, column := range columns {
		idx, err := FindColumnIndex(column, table.Header)
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
		if idx != expectedIdx[column] {
			t.Fatalf("expected index: %d, got=%d", expectedIdx[column], idx)
		}
	}

	idx, _ := FindColumnIndex("haha", table.Header)
	if idx != -1 {
		t.Fatalf("expected -1")
	}
}
