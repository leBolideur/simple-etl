package input

import "testing"

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
