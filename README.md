# Go Programming Learning Repository

[![Go Version](https://img.shields.io/badge/go-1.26.5-blue.svg)](https://go.dev/doc/go1.26)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Platform](https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-lightgray.svg)](https://go.dev/dl/)

A comprehensive Go (Golang) learning repository to help you get started with Go programming.

---

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Environment Setup](#environment-setup)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Debugging](#debugging)
- [Useful Resources](#useful-resources)

---

## ✅ Prerequisites

Before you begin, ensure you have met the following requirements:

- **Operating System**: Windows, macOS, or Linux
- **Go Version**: 1.26.5 or later (recommended)

---

## 📦 Installation

### Step 1: Download Go

Visit the official Go website and download the installer for your operating system:

- **Windows**: [go1.26.5.windows-amd64.msi](https://go.dev/dl/go1.26.5.windows-amd64.msi)
- **macOS**: [go1.26.5.darwin-amd64.pkg](https://go.dev/dl/go1.26.5.darwin-amd64.pkg)
- **Linux**: [go1.26.5.linux-amd64.tar.gz](https://go.dev/dl/go1.26.5.linux-amd64.tar.gz)

### Step 2: Run the Installer

Follow the installation wizard to complete the setup. The installer will automatically set up the `GOROOT` environment variable.

---

## 🔧 Environment Setup

### Set Environment Variables

#### Windows

1. Open **System Properties** → **Advanced** → **Environment Variables**
2. Add a new system variable:
   - **Variable name**: `GOPATH`
   - **Variable value**: `C:\Users\<YourUsername>\go` (or your preferred workspace)
3. Add `%GOPATH%\bin` to your `PATH` variable
4. Add `%GOROOT%\bin` to your `PATH` variable

#### macOS / Linux

Add the following to your shell configuration file (`~/.bashrc`, `~/.zshrc`, or `~/.profile`):

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

### Verify Installation

Open a terminal and run the following commands to verify your Go installation:

```bash
# Check Go version
go version

# Check environment variables
go env

# Print Go installation path
which go  # macOS/Linux
where go  # Windows
```

**Expected Output:**
```
go version go1.26.5 windows/amd64
```

---

## 📁 Project Structure

```
GO Programming/
├── hello_word/          # Main module
│   ├── go.mod           # Module definition
│   └── main.go          # Entry point
├── .gitignore           # Git ignore rules
└── README.md            # Project documentation
```

### Module Definition

The project uses Go modules with the module name `hello`:

```go
module hello

go 1.26.5
```

---

## 🚀 Usage

### Create a New Project

```bash
# Create a new directory for your project
mkdir myproject
cd myproject

# Initialize a new Go module
go mod init github.com/yourusername/myproject

# Create your first Go file
touch main.go
```

### Run the Project

Navigate to the `hello_word` directory and run the program:

```bash
cd hello_word

# Run the Go program
go run main.go
```

**Expected Output:**
```
TANVIR 25 Bangladesh true
```

### Build the Project

```bash
# Build for current platform
go build

# Build with output name
go build -o hello.exe

# Build for specific platform (cross-compilation)
GOOS=linux GOARCH=amd64 go build -o hello-linux
GOOS=darwin GOARCH=amd64 go build -o hello-macos
GOOS=windows GOARCH=amd64 go build -o hello-windows.exe
```

### Common Go Commands

| Command | Purpose |
|---------|---------|
| `go mod init myproject` | Initialize a new Go module |
| `go run main.go` | Run the program directly |
| `go build` | Compile into executable |
| `go build -o myapp.exe` | Compile with custom name |
| `go mod tidy` | Add missing dependencies |
| `go get <package>` | Download packages |
| `go test` | Run tests |
| `go vet` | Check for errors |

---

## 🐛 Debugging

### Using VS Code

1. Install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
2. Open your Go project in VS Code
3. Set breakpoints by clicking the gutter next to line numbers
4. Press `F5` to start debugging

### Using Delve (dlv)

```bash
# Install Delve debugger
go install github.com/go-delve/delve/cmd/dlv@latest

# Start debugging
dlv debug main.go

# Common Delve commands:
# continue (c) - continue execution
# next (n) - step to next line
# step (s) - step into function
# print (p) - print variable value
# break (b) - set breakpoint
```

---

## 📚 Useful Resources

- **Official Documentation**: [go.dev/doc](https://go.dev/doc/)
- **Go by Example**: [gobyexample.com](https://gobyexample.com/)
- **Effective Go**: [go.dev/doc/effective_go](https://go.dev/doc/effective_go)
- **Go Standard Library**: [pkg.go.dev](https://pkg.go.dev/)
- **Go Tutorial**: [go.dev/tour](https://go.dev/tour/)

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Happy coding in Go! 🚀