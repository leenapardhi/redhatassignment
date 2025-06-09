package csvutils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func AggregateCSV(headers []string, rows [][]string, columnName, operation string) (float64, error) {
	colIndex := -1

	for i, h := range headers {
		if strings.EqualFold(h, columnName) {
			colIndex = i
			break
		}
	}
	if colIndex == -1 {
		return 0, fmt.Errorf("column '%s' not found", columnName)
	}

	var sum, min, max float64
	var count int
	min = 1e18
	max = -1e18

	for _, row := range rows {
		if len(row) <= colIndex {
			continue
		}
		valStr := strings.TrimSpace(row[colIndex])
		if valStr == "" {
			continue
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			continue
		}
		sum += val
		count++

		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	if count == 0 {
		return 0, errors.New("no valid numeric values found")
	}

	switch strings.ToLower(operation) {
	case "sum":
		return sum, nil
	case "avg":
		return sum/float64(count), nil
	case "count":
		return float64(count), nil
	case "min":
		return min, nil
	case "max":
		return max, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}
