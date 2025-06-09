package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCSV(outputPath string, headers []string, data [][]string, summary []string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("cannot create output file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Write(headers)
	for _, row := range data {
		writer.Write(row)
	}

	writer.Write([]string{})

	for _, line := range summary {
		writer.Write([]string{line})
	}

	writer.Flush()
	return writer.Error()
}
