package filter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leBolideur/simple-etl/input"
	"github.com/leBolideur/simple-etl/utils"
)

type ColumnType int

const (
	ColumTypeInt ColumnType = iota
)

type intFilterFunc (func(int64, int64) bool)

var intFilterMap = map[byte]intFilterFunc{
	'>': func(value, filterValue int64) bool { return value > filterValue },
	'<': func(value, filterValue int64) bool { return value < filterValue },
}

type IntFilter struct {
	columnName  string
	ColumnType  ColumnType
	columnIndex int
	filterValue int64
	fn          intFilterFunc
}

func ApplyFilter(table *input.Table, filter string) error {
	filters, err := parseIntFilter(filter)
	if err != nil {
		return err
	}

	for i := range filters {
		colIndex, err := utils.FindColumnIndex(filters[i].columnName, table.Header)
		if err != nil {
			return err
		}

		filters[i].columnIndex = colIndex
	}

	filterMap := make(map[int]IntFilter)
	for _, filter := range filters {
		filterMap[filter.columnIndex] = filter
	}

	for i := range table.Rows {
		for j, cell := range table.Rows[i].Cells {
			if filter, ok := filterMap[j]; ok {
				var inferedValue = cell.InferedValue.(int64)
				if !filter.fn(inferedValue, filter.filterValue) {
					table.Rows[i].IsFiltered = true
				}
			}
		}
	}

	return nil
}

func parseIntFilter(filter string) ([]IntFilter, error) {
	filtersSplit := strings.Split(filter, ",")
	filters := make([]IntFilter, 0, len(filtersSplit))

	for _, mod := range filtersSplit {
		split := strings.Split(mod, ":")
		if len(split) < 2 {
			err := fmt.Errorf("filter expected format: <column_name>:<filter>\n")
			return nil, err
		}

		filterChar := split[1][0]
		filterFn, err := getFilterFunc(filterChar)
		if err != nil {
			return nil, err
		}

		filterIntValue, err := strconv.ParseInt(split[1][1:], 0, 64)
		if err != nil {
			err := fmt.Errorf("filter value is not an integer\n")
			return nil, err
		}
		intFilterStruct := IntFilter{
			columnName:  split[0],
			ColumnType:  ColumTypeInt,
			columnIndex: -1,
			filterValue: filterIntValue,
			fn:          filterFn,
		}

		filters = append(filters, intFilterStruct)
	}

	return filters, nil
}

func getFilterFunc(filterChar byte) (intFilterFunc, error) {
	fn, ok := intFilterMap[filterChar]
	if !ok {
		err := fmt.Errorf("no filter found for %q\n", filterChar)
		return nil, err
	}

	return fn, nil
}
