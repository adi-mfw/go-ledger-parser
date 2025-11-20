# Installation Guide

This guide explains how to download, install, and use the pre-built binary of `go-ledger-parser`.

## 📥 Downloading the Binary

### Option 1: From GitHub Releases (Recommended)

1. Go to the **Releases** page of this repository (click "Releases" in the right sidebar on GitHub)
2. Find the latest release (e.g., `v1.0.0`)
3. Download the appropriate binary for your platform:
   - **macOS (Apple Silicon)**: `ledger-parser-darwin-arm64`
   - **macOS (Intel)**: `ledger-parser-darwin-amd64` (if available)
   - **Linux**: `ledger-parser-linux-amd64` (if available)
   - **Windows**: `ledger-parser-windows-amd64.exe` (if available)

### Option 2: Using wget or curl

```bash
# macOS Apple Silicon (replace YOUR_USERNAME with your GitHub username)
wget https://github.com/YOUR_USERNAME/go-ledger-parser/releases/download/v1.0.0/ledger-parser-darwin-arm64

# Or with curl
curl -L -o ledger-parser-darwin-arm64 https://github.com/YOUR_USERNAME/go-ledger-parser/releases/download/v1.0.0/ledger-parser-darwin-arm64
```

**Note:** Replace `YOUR_USERNAME` with the actual GitHub username or organization name.

---

## 🔧 Installation

### Step 1: Make the Binary Executable

After downloading, make the binary executable:

```bash
chmod +x ledger-parser-darwin-arm64
```

### Step 2: Verify Installation

Test that the binary works:

```bash
./ledger-parser-darwin-arm64 --version
```

You should see the version number displayed.

### Step 3: (Optional) Install System-Wide

You can install the binary system-wide so you can run it from anywhere:

#### Option A: Install to `/usr/local/bin` (requires sudo)

```bash
sudo mv ledger-parser-darwin-arm64 /usr/local/bin/ledger-parser
```

Now you can run it from anywhere:
```bash
ledger-parser --input transactions.json --output report.json
```

#### Option B: Install to User Directory

```bash
# Create bin directory in your home folder
mkdir -p ~/bin

# Move binary there
mv ledger-parser-darwin-arm64 ~/bin/ledger-parser

# Add to PATH (add this line to ~/.zshrc or ~/.bash_profile)
export PATH="$HOME/bin:$PATH"

# Reload your shell configuration
source ~/.zshrc  # For zsh
# or
source ~/.bash_profile  # For bash
```

Now you can run `ledger-parser` from any directory.

---

## 🚀 Usage

### Basic Usage

```bash
./ledger-parser-darwin-arm64 --input transactions.json --output report.json
```

### Using Short Flags

```bash
./ledger-parser-darwin-arm64 -i transactions.json -o report.json
```

### With Verbose Output

```bash
./ledger-parser-darwin-arm64 -i transactions.json -o report.json -v
```

### If Installed System-Wide

```bash
ledger-parser --input transactions.json --output report.json
```

---

## 📝 Input File Format

Create a JSON file with your transaction data. Example `transactions.json`:

```json
[
  {
    "id": "txn001",
    "type": "CREDIT",
    "amount": 100.50,
    "date": "2024-01-15",
    "user_id": "user123"
  },
  {
    "id": "txn002",
    "type": "DEBIT",
    "amount": 50.25,
    "date": "2024-01-16",
    "user_id": "user123"
  },
  {
    "id": "txn003",
    "type": "CREDIT",
    "amount": 200.00,
    "date": "2024-01-17",
    "user_id": "user456"
  }
]
```

**Required fields:**
- `id`: Transaction identifier (string)
- `type`: "CREDIT" or "DEBIT" (string)
- `amount`: Positive number (float64)
- `date`: Date in "YYYY-MM-DD" format (string)
- `user_id`: User identifier (string)

---

## 📊 Output Format

The tool generates a JSON report with aggregated data per user:

```json
[
  {
    "user_id": "user123",
    "total_credit": 100.5,
    "total_debit": 50.25,
    "valid_count": 2,
    "invalid_count": 0
  },
  {
    "user_id": "user456",
    "total_credit": 200.0,
    "total_debit": 0,
    "valid_count": 1,
    "invalid_count": 0
  }
]
```

**Output fields:**
- `user_id`: User identifier
- `total_credit`: Sum of all valid credit transactions
- `total_debit`: Sum of all valid debit transactions
- `valid_count`: Number of valid transactions
- `invalid_count`: Number of invalid transactions

---

## ✅ Example Workflow

1. **Create your input file:**
   ```bash
   # Create transactions.json with your data
   nano transactions.json
   ```

2. **Run the parser:**
   ```bash
   ./ledger-parser-darwin-arm64 -i transactions.json -o report.json
   ```

3. **Check the output:**
   ```bash
   cat report.json
   ```

4. **View validation errors (if any):**
   The tool will print validation errors to the console for any invalid transactions.

---

## 🔍 Troubleshooting

### "Permission denied" error

Make sure the binary is executable:
```bash
chmod +x ledger-parser-darwin-arm64
```

### "Command not found" error

If you installed system-wide and get this error:
- Make sure you added the directory to your PATH
- Reload your shell: `source ~/.zshrc` or `source ~/.bash_profile`
- Verify the binary is in the correct location: `which ledger-parser`

### "required flag(s) not set" error

Make sure you provide both `--input` and `--output` flags:
```bash
./ledger-parser-darwin-arm64 --input input.json --output output.json
```

### "failed to read file" error

- Check that the input file exists
- Verify the file path is correct
- Ensure you have read permissions for the input file

### "failed to write file" error

- Check that you have write permissions in the output directory
- Ensure the output directory exists (the tool will create the file but not the directory)

---

## 🆘 Getting Help

Run the help command to see all available options:

```bash
./ledger-parser-darwin-arm64 --help
```

Or:

```bash
./ledger-parser-darwin-arm64 -h
```

---

## 📋 Command-Line Options

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--input` | `-i` | Input JSON file containing transactions | Yes |
| `--output` | `-o` | Output JSON file for the report | Yes |
| `--verbose` | `-v` | Enable verbose output | No |
| `--version` | | Show version information | No |
| `--help` | `-h` | Show help message | No |

---

## 🔄 Updating

To update to a newer version:

1. Download the new binary from the latest release
2. Replace the old binary with the new one
3. Make sure it's executable: `chmod +x ledger-parser-darwin-arm64`

---

## 📦 Requirements

- **No dependencies required!** The binary is self-contained
- **No Go installation needed** - the binary includes everything
- Compatible with:
  - macOS (Apple Silicon - M1/M2/M3)
  - macOS (Intel) - if binary available
  - Linux - if binary available
  - Windows - if binary available

---

## 🛠️ Building from Source

If you prefer to build from source instead of using the pre-built binary, see the main [README.md](README.md) for instructions.

---

*For more information, see the [README.md](README.md) file.*

