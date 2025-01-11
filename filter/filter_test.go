package filter

import "testing"

func TestParseIntFilter(t *testing.T) {
	input := "age:>50,year:<2020"

	filters, err := parseIntFilter(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(filters) != 2 {
		t.Fatalf("expected 2 filters, got=%d", len(filters))
	}

	expectedColumnName := []string{
		"age",
		"year",
	}
	expectedFilterValue := []int64{
		50,
		2020,
	}

	for i, filter := range filters {
		if filter.columnName != expectedColumnName[i] {
			t.Fatalf("expected columnName %q, got=%q", expectedColumnName[i], filter.columnName)
		}

		if filter.filterValue != expectedFilterValue[i] {
			t.Fatalf("expected filterValue %d, got=%d", expectedFilterValue[i], filter.filterValue)
		}
	}
}
