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
	fmt.Println(co.table.Header)
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

func WriteOutput(output string, table *input.Table) error {
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
