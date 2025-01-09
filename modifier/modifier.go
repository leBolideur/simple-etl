package modifier

import (
	"fmt"
	"strings"
)

type stringsModifierFunc (func(string) string)

var stringsModifiersMap = map[string]stringsModifierFunc{
	"uppercase": strings.ToUpper,
	"lowercase": strings.ToLower,
}

type Modifier struct {
	columnName  string
	columnIndex int
	fn          stringsModifierFunc
}

func ApplyModifier(rows [][]string, modifier string) ([]string, error) {
	modifiers, err := parseModifiers(modifier)
	if err != nil {
		return nil, err
	}
	header := rows[0]

	for i := range modifiers {
		colIndex, err := findColumnIndex(modifiers[i].columnName, header)
		if err != nil {
			return nil, err
		}

		modifiers[i].columnIndex = colIndex
	}

	modifierMap := make(map[int]stringsModifierFunc)
	for _, mod := range modifiers {
		modifierMap[mod.columnIndex] = mod.fn
	}

	var output = make([]string, 0, len(rows))
	output = append(output, strings.Join(header, ","))
	for _, line := range rows[1:] {
		var rowOutput strings.Builder

		for j, row := range line {
			if fn, ok := modifierMap[j]; ok {
				rowOutput.WriteString(fn(row))
			} else {
				rowOutput.WriteString(row)
			}

			if j < len(line)-1 {
				rowOutput.WriteString(",")
			}
		}
		output = append(output, rowOutput.String())
	}

	return output, nil
}

func findColumnIndex(colName string, header []string) (int, error) {
	colIndex := -1
	for i, column := range header {
		if column == colName {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		error := fmt.Errorf("no existing column with name: %s\n", colName)
		return colIndex, error
	}

	return colIndex, nil
}

func parseModifiers(modifier string) ([]Modifier, error) {
	modifiersSplit := strings.Split(modifier, ",")
	modifiers := make([]Modifier, 0, len(modifiersSplit))

	for _, mod := range modifiersSplit {
		split := strings.Split(mod, ":")
		if len(split) < 2 {
			err := fmt.Errorf("modifier expected format: <column_name>:<modifier>\n")
			return nil, err
		}

		modifierFn, err := getModifierFunc(split[1])
		if err != nil {
			return nil, err
		}

		modifierStruct := Modifier{
			columnName:  split[0],
			columnIndex: -1,
			fn:          modifierFn,
		}

		modifiers = append(modifiers, modifierStruct)
	}

	return modifiers, nil
}

func getModifierFunc(modifier string) (stringsModifierFunc, error) {
	fn, ok := stringsModifiersMap[modifier]
	if !ok {
		err := fmt.Errorf("no modifier found for %q\n", modifier)
		return nil, err
	}

	return fn, nil
}
