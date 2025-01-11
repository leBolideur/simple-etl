package modifier

import (
	"fmt"
	"strings"

	"github.com/leBolideur/simple-etl/input"
	"github.com/leBolideur/simple-etl/utils"
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

func ApplyModifier(table *input.Table, modifier string) error {
	modifiers, err := parseModifiers(modifier)
	if err != nil {
		return err
	}

	for i := range modifiers {
		colIndex, err := utils.FindColumnIndex(modifiers[i].columnName, table.Header)
		if err != nil {
			return err
		}

		modifiers[i].columnIndex = colIndex
	}

	modifierMap := make(map[int]Modifier)
	for _, mod := range modifiers {
		modifierMap[mod.columnIndex] = mod
	}

	for _, line := range table.Rows {
		for j, cell := range line.Cells {
			if modifier, ok := modifierMap[j]; ok {
				cell.RawValue = modifier.fn(cell.RawValue)
			}
		}
	}

	return nil
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
