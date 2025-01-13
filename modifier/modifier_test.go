package modifier

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

func TestApplyModifier(t *testing.T) {
	table := createTestTable()

	err := ApplyModifier(table, "first_name:uppercase,last_name:lowercase")
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	expectedUpper := []string{"JOHN", "JANE"}
	expectedLower := []string{"doe", "smith"}
	for i, row := range table.Rows {
		firstName := row.Cells[0].RawValue
		lastName := row.Cells[1].RawValue

		if firstName != expectedUpper[i] {
			t.Fatalf("expectedUpper: %q, got=%q", expectedUpper[i], firstName)
		}
		if lastName != expectedLower[i] {
			t.Fatalf("expectedUpper: %q, got=%q", expectedLower[i], lastName)
		}
	}
}

func TestParseModifier(t *testing.T) {
	input := "first_name:lowercase,username:uppercase,last_name:uppercase"

	modifiers, err := parseModifiers(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(modifiers) != 3 {
		t.Fatalf("expected 3 modifiers, got=%d", len(modifiers))
	}

	expectedColumnName := []string{
		"first_name",
		"username",
		"last_name",
	}

	for i, mod := range modifiers {
		if mod.columnName != expectedColumnName[i] {
			t.Fatalf("expected columnName %q, got=%q", expectedColumnName[i], mod.columnName)
		}
	}
}
