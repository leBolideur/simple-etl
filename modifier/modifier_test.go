package modifier

import "testing"

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
