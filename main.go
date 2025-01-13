package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/leBolideur/simple-etl/filter"
	"github.com/leBolideur/simple-etl/input"
	"github.com/leBolideur/simple-etl/modifier"
	"github.com/leBolideur/simple-etl/output"
)

func main() {
	args := os.Args
	if len(args) <= 2 {
		fmt.Println("Usage: ./simple-etl -f=<csv|json> --in=<filepath> --modifier=<column_name>:<modifier>")
		return
	}

	formatFlag := flag.String("f", "csv", "Select a format: CSV or JSON")
	inputFlag := flag.String("in", "", "Select an input file")
	outputFlag := flag.String("out", "cli", "Select an output")
	modifierFlag := flag.String("modifier", "", "Choose your modifier, eg: <column_name>:uppercase")
	filterFlag := flag.String("filter", "", "Choose your filter, eg: <column_name>:<10")
	flag.Parse()

	var table *input.Table
	if *formatFlag == "csv" {
		table_, err := input.CreateTableFromCSV(*inputFlag)
		table = table_
		if err != nil {
			fmt.Fprintf(os.Stderr, "readCSV err >> %s\n", err.Error())
			return
		}
	}

	if *filterFlag != "" {
		err := filter.ApplyFilter(table, *filterFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on applyModifier >> %s\n", err.Error())
		}
	}

	if *modifierFlag != "" {
		err := modifier.ApplyModifier(table, *modifierFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on applyModifier >> %s\n", err.Error())
		}
	}

	err := output.WriteOutput(*outputFlag, table)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error on writeOuput >> %s\n", err.Error())
	}
}
