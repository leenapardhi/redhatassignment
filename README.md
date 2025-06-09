# redhatassignment
# 📊 CSV File Processing Utility (Go)

This project is a lightweight and efficient command-line utility built in **Golang** to perform various operations on CSV files. It demonstrates real-world data processing techniques and is designed to be robust, extensible, and testable.

---

## ✨ Features

This utility supports the following operations on CSV files:

1. **Read and Display CSV Data**
   - Reads the CSV file and displays the first N rows in **comma-separated format**.
2. **Filter Rows Based on Criteria**
   - Supports filtering rows by:
     - Numeric conditions: `>`, `<`, `=`, `>=`, `<=`
     - String conditions: `contains`, `equals`
3. **Sort Rows by Column**
   - Sorts CSV rows by a specified column in **ascending** or **descending** order.
4. **Aggregate Data**
   - Aggregates numeric columns using functions like: `sum`, `average`, `min`, `max`.
5. **Count Valid Palindromes**
   - Detects and counts **palindromes** using only characters: `A`, `D`, `V`, `B`, and `N`.
6. **Write to a New CSV File**
   - Writes processed data (after filtering/sorting) into a **new CSV file**.
7. **Unit Tests**
   - Includes test coverage for each feature using Go’s `testing` package.

---

## 📂 Project Structure
csv-utility/
├── main.go # Main program entry
├── ecommerce_sales_data.csv # Sample input CSV file
├── csvutils/
│ ├── aggregate.go # Aggregate functions
| ├── countpalindrome.go # Palindrome functions
| ├── filter.go # Filter functions
| ├── reader.go # Read CSV File functions
| ├── sort.go # Sort functions
| ├── writerCSV.go # Write NewCSV File functions 
│ └── csvutils_test.go # Unit tests
└── README.md # Project documentation

🚀 Getting Started
**Prerequisites**
Go 1.20 or above installed
Git (to clone the repo)

**Clone the Repository**
git clone https://github.com/your-username/csv-utility.git
cd csv-utility

**Run the Application**
go run main.go
This will:
Read the CSV file
Display the first 3 rows
Perform filtering, sorting, aggregation
Count valid palindromes
Write processed results to output.csv

🧪 Run Tests
Run all tests with:
go test ./...
Expected Output:
ok      csv-utility/csvutils    0.XXXs

🛠️ Functions in Detail
  Functionality	                            Description
    ReadCSV	                          Reads a CSV file and returns headers and data
    PrintData	                        Displays the first N rows (headers + data) in comma-separated format
    FilterCSV	                        Filters rows using column name, operator, and condition
    SortCSV	                          Sorts rows based on a given column
    AggregateCSV	                    Aggregates values using sum, average, min, or max
    CountPalindrome	                  Counts valid palindromes with only A, D, V, B, and N
    WriteCSV	                        Saves headers, rows, and summary to a new file

💡 Edge Cases Handled
Empty rows or missing values
Invalid column names or indexes
Type mismatches during sorting or aggregation
Palindrome check case sensitivity and allowed character filtering

📬 Contact
Author: Leena Pardhi
Email: [leenapardhi@.com]
GitHub: https://github.com/leenapardhi
