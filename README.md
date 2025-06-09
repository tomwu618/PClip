
# PClip

[![Go Version][go-shield]][go-url]
[![MIT License][license-shield]][license-url]
[![CI Status][ci-shield]][ci-url]
[![Go Report Card][goreport-shield]][goreport-url]
[![Latest Release][release-shield]][release-url]

**PClip is a blazing-fast command-line tool that scans, aggregates, and copies source code to your clipboard, perfectly formatted in Markdown for AI-powered development.**

---

In the age of AI assistants like GPT-4 and Claude, developers constantly feed code snippets, files, or even entire project structures as context for prompts. This process is often manual, repetitive, and cumbersome. PClip streamlines this entire workflow into a single, simple command.

### User Story: The "Why"

> As a developer using AI assistants, I want to quickly select all relevant code files (e.g., all `.ts` and `.tsx` files in my `src/components` directory) and copy their contents to my clipboard in a clean, AI-friendly format, so that I can paste it directly into a prompt without manual copy-pasting or reformatting.

PClip was built to solve this exact problem.

## Key Features

- **Recursive File Scanning**: Instantly find files by extension in any directory.
- **AI-Optimized Formatting**: Wraps all code in Markdown blocks with language identifiers for optimal AI comprehension.
- **Clipboard Integration**: Sends the formatted output directly to your system clipboard.
- **Cross-Platform**: Built with Go to run natively on macOS, Windows, and Linux.
- **Simple & Fast**: A minimalist CLI with zero configuration and near-instant execution.

## Installation

### Prerequisites

- For Homebrew: A working installation of [Homebrew](https://brew.sh/).
- For Building from Source: A working [Go environment](https://go.dev/doc/install) (version 1.21+ recommended).

### 1. Homebrew (macOS & Linux)

This is the recommended method for macOS and Linux users.

```bash
# First, tap the repository
brew tap your-username/pclip

# Now, install pclip
brew install pclip
```

### 2. From GitHub Releases

You can download pre-compiled binaries for your specific OS and architecture directly from the Releases Page. Simply download the appropriate archive, extract the `pclip` executable, and place it in a directory within your system's PATH (e.g., `/usr/local/bin` or `C:\Windows\System32`).

### 3. Build from Source

If you prefer to build it yourself or are on a different platform:

```bash
# 1. Clone the repository
git clone https://github.com/your-username/pclip.git
cd pclip

# 2. Build the binary
go build -o pclip main.go

# 3. (Optional) Install it to your system path
# For Linux/macOS
sudo mv ./pclip /usr/local/bin/

# For Go developers, you can also use:
go install
```

## Usage

The command structure is simple:

```bash
pclip <ext1> [<ext2>...] [<directory>]
```

- Arguments in `<>` are required.
- Arguments in `[]` are optional.
- If no directory is specified, it defaults to the current directory (`.`).

### Example 1: Basic Use

Scan for all `.go` files in the current directory and its subdirectories.

```bash
pclip .go
```

### Example 2: Multiple Extensions & Specific Path

Scan for all TypeScript (`.ts`, `.tsx`) and CSS (`.scss`) files inside the `./src/app` directory.

```bash
pclip .ts .tsx .scss ./src/app
```

## Clipboard Output Format

After running a command, your clipboard will contain a clean, Markdown-formatted string like this, ready to be pasted into any AI chat prompt:

### `src/main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, PClip!")
}
```

### `src/utils/helpers.go`

```go
package utils

// Helper function
func Helper() bool {
    return true
}
```

## Development & Contributing

Contributions are welcome! This project is built with Go and uses a fully automated CI/CD pipeline with GitHub Actions.

### Local Development

- Clone the repository.
- Ensure you have Go installed.
- Run the application directly:

```bash
go run main.go .go ./your/test/dir
```

- Run tests:

```bash
go test -v ./...
```

All pull requests to the main branch are automatically built and tested. Releases are created automatically when a new `v*` tag is pushed.

## License

Distributed under the MIT License. See `LICENSE` for more information.

---

### **IMPORTANT: Next Steps**

Before you commit this `README.md` file, you need to replace all instances of `your-username` with your actual GitHub username. This is critical for the badge and installation URLs to work correctly.

For example, change:
`https://github.com/your-username/pclip`

to:
`https://github.com/john-doe/pclip`
