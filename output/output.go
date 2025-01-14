package output

import (
	"fmt"

	"github.com/leBolideur/simple-etl/input"
)

type Output interface {
	write()
}

type CLIOutput struct {
	table *input.Table
}

func (co CLIOutput) write() {
	headerCells := co.table.Header.Cells
	for i, col := range headerCells {
		fmt.Print(col.RawValue)
		if i < len(headerCells)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

	for _, row := range co.table.Rows {
		if row.IsFiltered {
			continue
		}
		for j, cell := range row.Cells {
			fmt.Print(cell.RawValue)
			if j < len(row.Cells)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

func WriteOutput(table *input.Table, output string) error {
	switch output {
	case "cli":
		cliOuput := CLIOutput{table: table}
		cliOuput.write()
	case "csv":
		// TODO
	default:
		err := fmt.Errorf("Invalid output interface")
		return err
	}

	return nil
}
