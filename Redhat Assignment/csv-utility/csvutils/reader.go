package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(filepath string) ([]string, [][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	if len(rows) == 0 {
		return nil, nil, fmt.Errorf("csv file is empty")
	}

	headers := rows[0]
	data := rows[1:]
	return headers, data, nil

}

func PrintData(headers []string, data [][]string, limit int) {
	fmt.Println(headers)
	for i, row := range data {
		if i >= limit {
			break
		}
		fmt.Println(row)
	}
}
