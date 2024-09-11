package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Count words from a chunk of text and handle word boundaries across chunks
func countWordsFromChunk(chunk string, inWord *bool) int {
	wordCount := 0
	for _, r := range chunk {
		if unicode.IsSpace(r) {
			*inWord = false
		} else if !*inWord {
			*inWord = true
			wordCount++
		}
	}
	return wordCount
}

// Count characters (runes) from a chunk of text, ensuring no split of UTF-8 runes across chunks
func countCharsFromChunk(chunk []byte, leftover []byte) (int, []byte) {
	combined := append(leftover, chunk...)

	// Check if the last rune is incomplete
	// We go backward from the end of the chunk to check if we encounter an incomplete UTF-8 character
	i := len(combined)
	for i > 0 && !utf8.FullRune(combined[i-1:]) {
		i--
	}

	// Return the count of characters and any leftover incomplete bytes
	return utf8.RuneCount(combined[:i]), combined[i:]
}

// Process input in chunks using bufio.Reader without loading the entire file into memory
func processFile(fileToOpen string, desiredFlags []string) error {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	if fileToOpen != "" {
		file, err := os.Open(fileToOpen)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	lineCount := 0
	wordCount := 0
	byteCount := 0
	charCount := 0

	buffer := make([]byte, 4096)  // 4KB buffer to read in chunks
	inWord := false               // Track if we're in the middle of a word
	leftover := []byte{}          // Leftover bytes for character handling

	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading input: %v", err)
		}
		if n == 0 {
			break
		}

		chunk := buffer[:n]

		// Count lines
		lineCount += strings.Count(string(chunk), "\n")

		// Count words, properly handling word boundaries between chunks
		wordCount += countWordsFromChunk(string(chunk), &inWord)

		// Count bytes
		byteCount += n

		// Count characters, properly handling UTF-8 rune boundaries between chunks
		charsInChunk, remaining := countCharsFromChunk(chunk, leftover)
		charCount += charsInChunk
		leftover = remaining

		if err == io.EOF {
			break
		}
	}

	// Handle any leftover bytes: if there are incomplete UTF-8 bytes left over at the end, ignore them
	if len(leftover) > 0 && utf8.FullRune(leftover) {
		charCount += utf8.RuneCount(leftover)
	}

	// Display the counts based on provided flags
	if contains(desiredFlags, "-l") {
		fmt.Printf("%d ", lineCount)
	}

	if contains(desiredFlags, "-w") {
		fmt.Printf("%d ", wordCount)
	}

	if contains(desiredFlags, "-c") {
		fmt.Printf("%d ", byteCount)
	}

	if contains(desiredFlags, "-m") {
		fmt.Printf("%d ", charCount)
	}

	// Default case if no flags are provided
	if !contains(desiredFlags, "-l") && !contains(desiredFlags, "-w") && !contains(desiredFlags, "-c") && !contains(desiredFlags, "-m") {
		fmt.Printf("%d %d %d ", lineCount, wordCount, byteCount)
	}

	fmt.Println(fileToOpen)

	return nil
}

// Helper function to check if a flag is present in the list
func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// Function to parse the command-line arguments and return desired flags and file to open
func parseArguments(args []string) ([]string, string, error) {
	var desiredFlags []string
	var fileToOpen string = ""

	for _, arg := range args[1:] {
		switch arg {
		case "-c", "-l", "-w", "-m":
			desiredFlags = append(desiredFlags, arg)
		default:
			if strings.Contains(arg, ".") {
				fileToOpen = arg
			} else {
				return nil, "", fmt.Errorf("ccwc: illegal option -- %s\nusage: ./ccwc [-Lclmw] [file ...]", arg)
			}
		}
	}

	return desiredFlags, fileToOpen, nil
}
