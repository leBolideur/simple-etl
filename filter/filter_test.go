package filter

import "testing"

func TestParseIntFilter(t *testing.T) {
	input := "age:>50,year:<2020,age:=50"

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
