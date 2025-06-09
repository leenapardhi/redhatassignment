package csvutils

import "strings"

func CountPalindrome(headers []string, data [][]string, column string) int {
	colIndex := -1
	for i, h := range headers {
		if strings.EqualFold(h, column) {
			colIndex = i
			break
		}
	}
	if colIndex == -1 {
		return 0
	}

	count := 0
	for _, row := range data {
		if colIndex < len(row) && isPalindrome(row[colIndex]) {
			count++
		}
	}
	return count
}

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
