// main.go
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
)

// main is the application's entry point.
func main() {
	// 1. Parse command-line arguments.
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		return
	}

	extensions, path := parseArgs(args)
	if len(extensions) == 0 {
		fmt.Println("Error: No file extensions specified.")
		printUsage()
		return
	}

	// 2. Find matching files.
	files, err := findFiles(path, extensions)
	if err != nil {
		log.Fatalf("Error: Failed to find files: %v", err)
	}

	if len(files) == 0 {
		fmt.Println("No matching files found in the directory.")
		return
	}

	// 3. Read file contents and format as Markdown.
	markdownContent, err := formatContentAsMarkdown(path, files)
	if err != nil {
		log.Fatalf("Error: Failed to format file content: %v", err)
	}

	// 4. Copy the content to the clipboard.
	if err := clipboard.WriteAll(markdownContent); err != nil {
		log.Fatalf("Error: Failed to copy to clipboard: %v", err)
	}

	// 5. Print a success message.
	fmt.Printf("✅ Success! Content of %d file(s) formatted and copied to clipboard.\n", len(files))
}

// printUsage prints the command-line usage instructions.
func printUsage() {
	fmt.Println(`
PClip: A clipboard tool for preparing code for AI prompts.

Usage:
  pclip <ext1> [<ext2>...] [<directory>]

Arguments:
  <ext>:       Extension must start with a '.', e.g., .go, .js, .ts
  <directory>: Optional. The directory to scan. Defaults to the current directory if not provided.

Examples:
  # Scan for all .go files in the current directory
  pclip .go

  # Scan for all .ts and .tsx files in the ./src directory
  pclip .ts .tsx ./src
	`)
}

// parseArgs parses command-line arguments to separate extensions and the target path.
func parseArgs(args []string) (extensions []string, path string) {
	path = "." // Default path is the current directory.

	// Check if the last argument is a directory.
	if len(args) > 0 {
		lastArg := args[len(args)-1]
		if info, err := os.Stat(lastArg); err == nil && info.IsDir() {
			path = lastArg
			args = args[:len(args)-1] // The rest are extensions.
		}
	}

	// Filter all arguments that start with '.' as extensions.
	for _, arg := range args {
		if strings.HasPrefix(arg, ".") {
			extensions = append(extensions, arg)
		}
	}

	return extensions, path
}

// findFiles recursively finds files with the given extensions in the specified root directory.
func findFiles(root string, extensions []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Handle potential access errors.
		}
		// Skip directories themselves and hidden files/folders.
		if d.IsDir() || strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		for _, ext := range extensions {
			if strings.HasSuffix(path, ext) {
				files = append(files, path)
				break
			}
		}
		return nil
	})
	return files, err
}

// formatContentAsMarkdown reads file contents and formats them into Markdown code blocks.
func formatContentAsMarkdown(rootPath string, files []string) (string, error) {
	var builder strings.Builder

	for i, file := range files {
		// Get the relative path for cleaner output.
		relativePath, err := filepath.Rel(rootPath, file)
		if err != nil {
			relativePath = file // Fallback to the full path if Rel fails.
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %w", file, err)
		}

		lang := getLanguageFromExtension(filepath.Ext(file))

		// Write the content in Markdown format.
		builder.WriteString(fmt.Sprintf("`%s`\n\n", relativePath)) // File path as a title.
		builder.WriteString(fmt.Sprintf("```%s\n", lang))
		builder.Write(content)
		builder.WriteString("\n```")

		// Add two newlines to separate file blocks.
		if i < len(files)-1 {
			builder.WriteString("\n\n")
		}
	}

	return builder.String(), nil
}

// getLanguageFromExtension returns a Markdown language identifier based on the file extension.
func getLanguageFromExtension(ext string) string {
	// Remove the leading dot.
	ext = strings.TrimPrefix(ext, ".")
	// Simple mapping for common languages.
	switch ext {
	case "go":
		return "go"
	case "js":
		return "javascript"
	case "ts", "tsx":
		return "typescript"
	case "py":
		return "python"
	case "java":
		return "java"
	case "c":
		return "c"
	case "cpp":
		return "cpp"
	case "cs":
		return "csharp"
	case "html":
		return "html"
	case "css":
		return "css"
	case "scss":
		return "scss"
	case "sh", "bash":
		return "bash"
	case "rb":
		return "ruby"
	case "rs":
		return "rust"
	case "md":
		return "markdown"
	case "json":
		return "json"
	case "yaml", "yml":
		return "yaml"
	case "sql":
		return "sql"
	default:
		return "" // No specific language identifier.
	}
}
