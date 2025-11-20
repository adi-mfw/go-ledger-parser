# go-ledger-parser
go-ledger-parser: Go Fundamentals &amp; File I/O Utility

This repository contains Proof of Concept 1 (POC 1) of the Accelerated Go Ramp-up.

🎯 Goal

The primary objective of this utility is to demonstrate mastery of core Go fundamentals, including strong typing, struct definition, idiomatic error handling, and file input/output (I/O).

💻 Deliverable

A command-line utility that simulates reading raw transaction data (mock CSV/JSON files), validates critical fields (e.g., amount > 0), processes the data into Go structs, and writes a summarized report to a new output file.

Key Topics Demonstrated

Go Basics: Variables, Functions, Packages.

Data Modeling: Structs, Methods, and a simple Interface for abstraction.

I/O Operations: Using the os, io, and encoding/json packages.

Error Handling: Consistent and idiomatic use of if err != nil.

## 📦 Installation

**Want to use the pre-built binary?** See the [Installation Guide](INSTALLATION.md) for detailed instructions on downloading and using the binary from GitHub Releases.

**Want to build from source?** Continue reading below.

## 🚀 How to Run

### Prerequisites
- Go 1.16 or later installed

### Quick Start

1. **Run directly with `go run`:**
   ```bash
   go run main.go --input sample_input.json --output output.json
   ```

2. **Or build and run:**
   ```bash
   # Build the executable
   go build -o ledger-parser main.go
   
   # Run the executable
   ./ledger-parser --input sample_input.json --output output.json
   ```

### Usage
```bash
./ledger-parser --input <input_json_file> --output <output_json_file>
```

You can also use short flags:
```bash
./ledger-parser -i <input_json_file> -o <output_json_file>
```

### Example
```bash
go run main.go --input sample_input.json --output report.json
```

Or with short flags:
```bash
go run main.go -i sample_input.json -o report.json
```

This will:
- Read transactions from `sample_input.json`
- Validate each transaction (amount > 0, valid date format)
- Print validation errors for invalid transactions
- Generate an aggregated report by user
- Write the report to `report.json`

### Sample Input Format
See `sample_input.json` for an example. Each transaction should have:
- `id`: Transaction identifier
- `type`: "CREDIT" or "DEBIT"
- `amount`: Positive number
- `date`: Date in "YYYY-MM-DD" format
- `user_id`: User identifier

### Output Format
The output JSON contains aggregated data per user:
- `user_id`: User identifier
- `total_credit`: Sum of all valid credit transactions
- `total_debit`: Sum of all valid debit transactions
- `valid_count`: Number of valid transactions
- `invalid_count`: Number of invalid transactions
