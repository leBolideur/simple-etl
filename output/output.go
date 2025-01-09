package output

import "fmt"

type Output interface {
	write()
}

type CLIOutput struct {
	rows []string
}

func (co CLIOutput) write() {
	for _, row := range co.rows {
		fmt.Println(row)
	}
}

func WriteOutput(output string, rows []string) error {
	switch output {
	case "cli":
		cliOuput := CLIOutput{rows: rows}
		cliOuput.write()
	case "csv":
		// TODO
	default:
		err := fmt.Errorf("Invalid output interface")
		return err
	}

	return nil
}
