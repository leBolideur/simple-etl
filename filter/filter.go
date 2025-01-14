package filter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leBolideur/simple-etl/input"
	"github.com/leBolideur/simple-etl/utils"
)

func ApplyFilter(table *input.Table, filter string) error {
	filters, err := parseFilters(filter)
	if err != nil {
		return err
	}

	for i := range filters {
		colIndex, err := utils.FindColumnIndex(filters[i].getColumnName(), table.Header)
		if err != nil {
			return err
		}

		filters[i].setColumnIndex(colIndex)
	}

	filterMap := make(map[int][]IFilter)
	for _, filter := range filters {
		filterMap[filter.getColumnIndex()] = append(filterMap[filter.getColumnIndex()], filter)
	}

	for _, row := range table.Rows {
		for colIdx, cell := range row.Cells {
			if filters, ok := filterMap[colIdx]; ok {
				for _, filter := range filters {
					if filter.getColumnType() != ColumnType(cell.Type) {
						err := fmt.Errorf("column %q of type %q are not compatible with %q", filter.getColumnName(), ColumnType(cell.Type), filter.getColumnType())
						return err
					}
					if !row.IsFiltered && !filter.apply(cell.InferedValue) {
						row.IsFiltered = true
					}
				}
			}
		}
	}

	return nil
}

func parseFilters(filter string) ([]IFilter, error) {
	filtersSplit := strings.Split(filter, ",")
	filters := make([]IFilter, 0, len(filtersSplit))

	for _, filter := range filtersSplit {
		split := strings.Split(filter, ":")
		if len(split) < 3 {
			err := fmt.Errorf("filter expected format: <column_name>:<filter>:<filter_value>\n")
			return nil, err
		}

		columnName, filterStr, filterValue := split[0], split[1], split[2]
		if strings.HasPrefix(filterStr, "len_") {
			lenFilter, err := parseStrLenFilter(columnName, filterStr, filterValue)
			if err != nil {
				return nil, err
			}

			filters = append(filters, lenFilter)
		} else {
			filter, err := inferFilterType(columnName, filterStr, filterValue)
			if err != nil {
				return nil, err
			}

			filters = append(filters, filter)
		}

	}

	return filters, nil
}

func parseStrLenFilter(columnName, filterStr, filterValue string) (IFilter, error) {
	lenValue, err := strconv.ParseInt(filterValue, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("filter value for len_* must be an integer")
	}

	filterFn, err := getStrLenFilterFunc(filterStr)
	if err != nil {
		return nil, err
	}

	lenFilter := &LenFilter{
		columnName:  columnName,
		columnType:  ColumnTypeString,
		columnIndex: -1,
		filterStr:   filterStr,
		filterValue: int(lenValue),
		fn:          filterFn,
	}

	return lenFilter, nil
}

func inferFilterType(columnName, filterStr, filterValue string) (IFilter, error) {
	filterIntValue, err := strconv.ParseInt(filterValue, 0, 64)
	if err == nil {
		filterFn, err := getIntFilterFunc(filterStr)
		if err != nil {
			return nil, err
		}

		intFilterStruct := &GenericFilter[int64]{
			columnName:  columnName,
			columnType:  ColumnTypeInt,
			columnIndex: -1,
			filterStr:   filterStr,
			filterValue: filterIntValue,
			fn:          filterFn,
		}

		return intFilterStruct, nil
	}

	filterBoolValue, err := strconv.ParseBool(filterValue)
	if err == nil {
		filterFn, err := getBoolFilterFunc(filterStr)
		if err != nil {
			return nil, err
		}

		boolFilterStruct := &GenericFilter[bool]{
			columnName:  columnName,
			columnType:  ColumnTypeBool,
			columnIndex: -1,
			filterStr:   filterStr,
			filterValue: filterBoolValue,
			fn:          filterFn,
		}

		return boolFilterStruct, nil
	}

	return nil, fmt.Errorf("unsupported filter value: %q", filterValue)
}

func getIntFilterFunc(filterStr string) (intFilterFunc, error) {
	fn, ok := intFilterMap[filterStr]
	if !ok {
		err := fmt.Errorf("no filter found for %q\n", filterStr)
		return nil, err
	}

	return fn, nil
}

func getBoolFilterFunc(filterStr string) (boolFilterFunc, error) {
	fn, ok := boolFilterMap[filterStr]
	if !ok {
		err := fmt.Errorf("no filter found for %q\n", filterStr)
		return nil, err
	}

	return fn, nil
}

func getStrLenFilterFunc(filterStr string) (strLenFilterFunc, error) {
	fn, ok := strLenFilterMap[filterStr]
	if !ok {
		err := fmt.Errorf("no filter found for %q\n", filterStr)
		return nil, err
	}

	return fn, nil
}
