---

# GopherCalc

GopherCalc is a command-line calculator application built in Go using the Cobra library. It provides basic arithmetic operations such as addition, subtraction, multiplication, and division.

## Installation

To use GopherCalc, you need to build the binary executable using the Go build command:

```bash
go build .
```

This will generate an executable file named `GopherCalc` in the current directory. You can then run this executable to use the calculator.

## Usage

GopherCalc provides several commands for performing arithmetic operations. Here are the available commands:

- `Addition`: Performs the addition operation for two numbers.
- `Subtraction`: Performs the subtraction operation for two numbers.
- `Multiplication`: Performs the multiplication operation for two numbers.
- `Division`: Performs the division operation for two numbers.

### Command Syntax

```bash
GopherCalc [command] [arguments]
```

### Example Usage

```bash
# Addition
GopherCalc Addition 10 5

# Subtraction
GopherCalc Subtraction 10 5

# Multiplication
GopherCalc Multiplication 10 5

# Division
GopherCalc Division 10 5
```

### Flags

- `-h, --help`: Displays help information about GopherCalc and its commands.
- `-t, --toggle`: Toggles a feature (example flag).

### Command Usage

For detailed usage instructions for each command, you can use the `--help` flag:

```bash
GopherCalc [command] --help
```

## License

This project is licensed under the [MIT License](LICENSE).

---
