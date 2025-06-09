package csvutils

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func SortCSV(headers []string, data [][]string, columnName string, ascending bool) ([][]string, error) {

	colIndex := -1
	for i, h := range headers {
		if strings.EqualFold(h, columnName) {
			colIndex = i
			break
		}
	}
	if colIndex == -1 {
		return nil, errors.New("column not found")
	}

	sorted := make([][]string, len(data))
	for i, row := range data {
		sorted[i] = append([]string(nil), row...)
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		a := sorted[i][colIndex]
		b := sorted[j][colIndex]
		aNum, errA := strconv.ParseFloat(a, 64)
		bNum, errB := strconv.ParseFloat(b, 64)

		if errA == nil && errB == nil {
			if ascending {
				return aNum < bNum
			}
			return aNum > bNum
		}

		if ascending {
			return strings.ToLower(a) < strings.ToLower(b)
		}
		return strings.ToLower(a) > strings.ToLower(b)
	})

	return sorted, nil

}
