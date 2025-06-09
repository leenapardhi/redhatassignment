package csvutils

import (
	"os"
	"strconv"
	"testing"
)

func TestReadCSV(t *testing.T) {
	headers, data, err := ReadCSV("../ecommerce_sales_data.csv")
	if err != nil {
		t.Fatalf("Failed to read CSV: %v", err)
	}
	if len(headers) == 0 || len(data) == 0 {
		t.Error("Headers or data should not be empty")
	}
}

func TestFilterCSV(t *testing.T) {
	headers, filtered, err := FilterCSV("../ecommerce_sales_data.csv", "Product", "contains", "Laptop")
	if err != nil {
		t.Fatalf("FilterCSV failed: %v", err)
	}
	if len(filtered) == 0 {
		t.Error("Expected some rows after filtering")
	}
	_ = headers
}

func TestSortCSV(t *testing.T) {
	headers, data, _ := ReadCSV("../ecommerce_sales_data.csv")
	sorted, err := SortCSV(headers, data, "Quantity", true)
	if err != nil {
		t.Fatalf("SortCSV failed: %v", err)
	}
	colIndex := -1
	for i, h := range headers {
		if h == "Quantity" {
			colIndex = i
			break
		}
	}
	if colIndex == -1 {
		t.Fatalf("Quantity column not found in headers")
	}

	for i := 1; i < len(sorted); i++ {
		prevVal, err1 := strconv.Atoi(sorted[i-1][colIndex])
		currVal, err2 := strconv.Atoi(sorted[i][colIndex])
		if err1 != nil || err2 != nil {
			t.Errorf("Failed to convert string to int at row %d or %d", i-1, i)
			continue
		}
		if prevVal > currVal {
			t.Errorf("Sort order incorrect at index %d: %d > %d", i, prevVal, currVal)
		}
	}
}

func TestAggregateCSV(t *testing.T) {
	headers, data, _ := ReadCSV("../ecommerce_sales_data.csv")
	_, err := AggregateCSV(headers, data, "UnitPrice", "sum")
	if err != nil {
		t.Errorf("AggregateCSV failed: %v", err)
	}
}

func TestCountPalindromes(t *testing.T) {
	headers, data, _ := ReadCSV("../ecommerce_sales_data.csv")
	count := CountPalindrome(headers, data, "CustomerName")
	if count < 0 {
		t.Error("Palindrome count should be >= 0")
	}
}

func TestWriteCSV(t *testing.T) {
	headers := []string{"Col1", "Col2"}
	rows := [][]string{{"Val1", "Val2"}, {"Test1", "Test2"}}
	filename := "test_output.csv"
	summary := []string{"Test Summary"}
	err := WriteCSV(filename, headers, rows, summary)
	if err != nil {
		t.Errorf("WriteCSV failed: %v", err)
	}
	defer os.Remove(filename)
}
