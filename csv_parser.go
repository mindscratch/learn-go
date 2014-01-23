package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func usage(programName string) {
	fmt.Println("usage:", programName, " /path/to/file.csv")
	os.Exit(1)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		usage(args[0])
	}

	filePath := args[1]
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Println("problem opening file", err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	reader.TrailingComma = true // allow the last field in any row to be empty

	for {
		fields, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println("FIELDS", fields)
	}
}
