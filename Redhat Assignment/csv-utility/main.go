package main

import (
	"csv-utility/csvutils"
	"fmt"
	"log"
	"strings"
)

func isValidColumn(headers []string, col string) bool {
	for _, h := range headers {
		if strings.EqualFold(h, col) {
			return true
		}
	}
	return false
}

func main() {
	filePath := "ecommerce_sales_data.csv"

	fmt.Println("---Reading CSV File---")
	headers, data, err := csvutils.ReadCSV(filePath)
	if err != nil {
		log.Fatal("Failed to read CSV", err)
	}
	fmt.Println("---First 3 rows---")
	// Print header as CSV line
	fmt.Println(strings.Join(headers, ","))

	// Print first 3 rows or fewer if data less than 3
	n := 3
	if len(data) < 3 {
		n = len(data)
	}
	for i := 0; i < n; i++ {
		fmt.Println(strings.Join(data[i], ","))
	}

	var summary []string
	summary = append(summary, "--- User Input and Output summary---")

	var columnName string
	for {
		fmt.Print("Enter column name to filter on: ")
		fmt.Scanln(&columnName)
		if isValidColumn(headers, columnName) {
			break
		}
		fmt.Printf("Column '%s' not found. Please try again.\n", columnName)
	}

	fmt.Print("Enter filter operator(==, !=, >, <, >=, <=, contains): ")
	var operator string
	fmt.Scanln(&operator)

	fmt.Print("Enter value to compare: ")
	var filterValue string
	fmt.Scanln(&filterValue)

	headers, filteredData, err := csvutils.FilterCSV(filePath, columnName, operator, filterValue)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("---Filtered Rows---")
	csvutils.PrintData(headers, filteredData, 10)
	summary = append(summary, fmt.Sprintf("Filtered by %s %s %s", columnName, operator, filterValue))

	// Store filtered separately
	finalData := filteredData

	//2. Sorting
	fmt.Print("Do you want to sort the CSV? (yes/no): ")
	var sortChoice string
	fmt.Scanln(&sortChoice)
	if strings.ToLower(sortChoice) == "yes" {
		fmt.Print("Enter column name to sort by: ")
		var sortColumn string
		fmt.Scanln(&sortColumn)

		fmt.Print("Enter order (asc/desc): ")
		var order string
		fmt.Scanln(&order)

		ascending := strings.ToLower(order) == "asc"

		sortedData, err := csvutils.SortCSV(headers, filteredData, sortColumn, ascending)
		if err != nil {
			fmt.Println("Error sorting CSV:", err)
		} else {
			fmt.Println("---Sorted Rows---")
			csvutils.PrintData(headers, sortedData, 10)
			summary = append(summary, fmt.Sprintf("Sorted by %s (%s)", sortColumn, order))
			finalData = sortedData
		}
	}

	// 3.Aggregate
	fmt.Print("Do you want to perform aggregation? (yes/no): ")
	var aggChoice string
	fmt.Scanln(&aggChoice)

	if strings.ToLower(aggChoice) == "yes" {
		fmt.Print("Enter column name for aggregation: ")
		var aggColumn string
		fmt.Scanln(&aggColumn)

		fmt.Print("Enter operation (sum/avg/count/min/max): ")
		var operation string
		fmt.Scanln(&operation)

		result, err := csvutils.AggregateCSV(headers, finalData, aggColumn, operation)
		if err != nil {
			fmt.Println("Aggregation Error:", err)
		} else {
			fmt.Printf("%s of '%s': %.2f\n", strings.Title(operation), aggColumn, result)
			summary = append(summary, fmt.Sprintf("Aggregation: %s of %s = %.2f", operation, aggColumn, result))
		}
	}

	// 4.Creating New CSV File
	fmt.Print("Do you want to export the result to a CSV file? (yes/no): ")
	var exportChoice string
	fmt.Scanln(&exportChoice)

	if strings.ToLower(exportChoice) == "yes" {
		fmt.Print("Enter the output fileName (example: result.csv): ")
		var outputFile string
		fmt.Scanln(&outputFile)

		err = csvutils.WriteCSV(outputFile, headers, filteredData, summary)
		if err != nil {
			fmt.Println("Error exporting csv: ", err)
		} else {
			fmt.Println("Data exported successfully to ", outputFile)
		}
	}

	// 5.Count Palindromes
	fmt.Print("Do you want to count palindromes in any column? (yes/no): ")
	var palChoice string
	fmt.Scanln(&palChoice)
	if strings.ToLower(palChoice) == "yes" {
		fmt.Print("Enter column name to check palindromes: ")
		var palColumn string
		fmt.Scanln(&palColumn)

		palCount := csvutils.CountPalindrome(headers, finalData, palColumn)
		fmt.Printf("Total Palindromes Found in '%s': %d\n", palColumn, palCount)
	}
}
