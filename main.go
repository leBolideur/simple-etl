package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

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
	flag.Parse()

	var rows [][]string
	if *formatFlag == "csv" {
		content, err := readCSV(*inputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readCSV err >> %s\n", err.Error())
			return
		}

		rows = content
	}

	var modifiedRows []string
	if *modifierFlag != "" {
		output, err := modifier.ApplyModifier(rows, *modifierFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on applyModifier >> %s\n", err.Error())
		}

		modifiedRows = output
	}

	err := output.WriteOutput(*outputFlag, modifiedRows)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error on writeOuput >> %s\n", err.Error())
	}
}

func readCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return content, nil
}
