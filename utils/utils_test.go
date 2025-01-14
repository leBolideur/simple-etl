package utils

import (
	"strings"
	"testing"

	"github.com/leBolideur/simple-etl/input"
)

func TestFindColumnIndex(t *testing.T) {
	inputStr := `first_name,last_name,age,active
			  John,Doe,30,true
			  Jane,Smith,45,false`

	reader := strings.NewReader(inputStr)
	table, err := input.CreateTableFromCSV(reader)
	if err != nil {
		t.Fatalf("Error on create table > %s", err.Error())
	}

	expectedIdx := map[string]int{
		"first_name": 0,
		"last_name":  1,
		"age":        2,
		"active":     3,
	}
	columns := []string{"first_name", "last_name", "age", "active"}
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
