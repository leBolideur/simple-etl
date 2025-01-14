package filter

type ColumnType string

const (
	ColumnTypeInt  ColumnType = "int"
	ColumnTypeBool            = "bool"
)

type intFilterFunc (func(int64, int64) bool)

var intFilterMap = map[string]intFilterFunc{
	">":  func(value, filterValue int64) bool { return value > filterValue },
	"gt": func(value, filterValue int64) bool { return value > filterValue },

	"<":  func(value, filterValue int64) bool { return value < filterValue },
	"lt": func(value, filterValue int64) bool { return value < filterValue },

	">=":  func(value, filterValue int64) bool { return value >= filterValue },
	"gte": func(value, filterValue int64) bool { return value >= filterValue },

	"<=":  func(value, filterValue int64) bool { return value <= filterValue },
	"lte": func(value, filterValue int64) bool { return value <= filterValue },

	"=":  func(value, filterValue int64) bool { return value == filterValue },
	"eq": func(value, filterValue int64) bool { return value == filterValue },

	"!=": func(value, filterValue int64) bool { return value != filterValue },
	"ne": func(value, filterValue int64) bool { return value != filterValue },
}

type boolFilterFunc (func(bool, bool) bool)

var boolFilterMap = map[string]boolFilterFunc{
	"=":  func(value, filterValue bool) bool { return value == filterValue },
	"eq": func(value, filterValue bool) bool { return value == filterValue },

	"!=": func(value, filterValue bool) bool { return value != filterValue },
	"ne": func(value, filterValue bool) bool { return value != filterValue },
}

type IFilter interface {
	apply(value any) bool
	getColumnName() string
	getColumnIndex() int
	setColumnIndex(int)
	getColumnType() ColumnType
}

type GenericFilter[T any] struct {
	columnName  string
	columnType  ColumnType
	columnIndex int

	filterStr   string
	filterValue T
	fn          func(T, T) bool
}

func (gf GenericFilter[T]) apply(value any) bool {
	v, ok := value.(T)
	if !ok {
		return false
	}

	return gf.fn(v, gf.filterValue)
}
func (gf GenericFilter[T]) getColumnName() string     { return gf.columnName }
func (gf GenericFilter[T]) getColumnIndex() int       { return gf.columnIndex }
func (gf *GenericFilter[T]) setColumnIndex(idx int)   { gf.columnIndex = idx }
func (gf GenericFilter[T]) getFilterValue() T         { return gf.filterValue }
func (gf GenericFilter[T]) getColumnType() ColumnType { return gf.columnType }
