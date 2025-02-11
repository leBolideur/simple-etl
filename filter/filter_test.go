package filter

import (
	"strings"
	"testing"

	"github.com/leBolideur/simple-etl/input"
)

func TestApplyFilter(t *testing.T) {
	inputStr := `first_name,last_name,age,active
			  John,Doe,30,true
			  Jane,Smith,45,false`

	reader := strings.NewReader(inputStr)
	table, err := input.CreateTableFromCSV(reader)

	if err != nil {
		t.Fatalf("Error on create table > %s", err.Error())
	}

	err = ApplyFilter(table, "age:eq:30,active:eq:true")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if table.Rows[0].IsFiltered != false {
		t.Fatalf("expected row[0].IsFiltered: %t, got=%t", false, table.Rows[0].IsFiltered)
	}
	if table.Rows[1].IsFiltered != true {
		t.Fatalf("expected row[1].IsFiltered: %t, got=%t", true, table.Rows[1].IsFiltered)
	}
}

func TestParseFilter(t *testing.T) {
	input := "age:gt:50,year:lt:2020,active:eq:true,first_name:len_eq:5"

	filters, err := parseFilters(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(filters) != 4 {
		t.Fatalf("expected 4 filters, got=%d", len(filters))
	}

	expected := []struct {
		columnName string
		columnType ColumnType
	}{
		{"age", ColumnTypeInt},
		{"year", ColumnTypeInt},
		{"active", ColumnTypeBool},
		{"first_name", ColumnTypeString},
	}

	for i, filter := range filters {
		if filter.getColumnName() != expected[i].columnName {
			t.Fatalf("expected columnName %q, got=%q", expected[i].columnName, filter.getColumnName())
		}

		if filter.getColumnType() != expected[i].columnType {
			t.Fatalf("expected filter type %v, got=%v", expected[i].columnType, filter.getColumnType())
		}
	}
}
