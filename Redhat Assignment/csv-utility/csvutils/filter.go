package csvutils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func FilterCSV(filePath, column, operator, value string) ([]string, [][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot open the file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot read headers: %v", err)
	}

	colIndex := -1
	for i, h := range header {
		if strings.EqualFold(h, column) {
			colIndex = i
			break
		}
	}
	if colIndex == -1 {
		return nil, nil, fmt.Errorf("column '%s' not found in CSV", column)
	}

	var filtered [][]string

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(row) != len(header) {
			continue
		}
		cell := row[colIndex]
		match, err := compare(cell, operator, value)
		if err != nil {
			continue
		}
		if match {
			filtered = append(filtered, row)
		}
	}
	return header, filtered, nil
}

func compare(cell, op, val string) (bool, error) {
	switch op {
	case "contains":
		return strings.Contains(strings.ToLower(cell), strings.ToLower(val)), nil

	case "==":
		return strings.EqualFold(cell, val), nil

	case "!=":
		return strings.EqualFold(cell, val), nil

	case ">", "<", ">=", "<=":
		num1, err1 := strconv.ParseFloat(cell, 64)
		num2, err2 := strconv.ParseFloat(val, 64)
		if err1 != nil || err2 != nil {
			return false, errors.New("non-numeric comparison on numeric operator")
		}

		switch op {
		case ">":
			return num1 > num2, nil
		case "<":
			return num1 < num2, nil
		case ">=":
			return num1 >= num2, nil
		case "<=":
			return num1 <= num2, nil
		}

	}

	return false, fmt.Errorf("unsupported operator: %s", op)
}
