package main

import (
	A "asciiart-fs/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Color codes for ANSI terminal
var colorCodes = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"reset":   "\033[0m",
}

var (
	_map  = make(map[int][8]string)
	lines = [8]string{}
)

// Insert value into the ASCII map
func InsertValue(scanner *bufio.Scanner) [8]string {
	ArtValue := [8]string{}
	for cp := 0; cp < 8 && scanner.Scan(); cp++ {
		ArtValue[cp] = scanner.Text()
	}
	scanner.Scan() // Skip extra newline if needed
	return ArtValue
}

// Check if argument is valid (printable characters only)
func IsValidArg(arg string) bool {
	if arg == "" {
		return false
	}
	for _, val := range arg {
		if val <= 31 || val >= 127 {
			return false
		}
	}
	return true
}

// Initialize the ASCII map from banner files
func InitMap(str string) {
	var file *os.File
	var err error
	switch str {
	case "standard":
		file, err = os.Open("Banners/standard.txt")
	case "shadow":
		file, err = os.Open("Banners/shadow.txt")
	case "thinkertoy":
		file, err = os.Open("Banners/thinkertoy.txt")
	default:
		os.Stderr.WriteString("Err: Invalid [BANNER]: " + str + "\n")
		os.Exit(0)
	}

	if err != nil {
		os.Stderr.WriteString("Err: " + err.Error() + "\n")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(file)
	// Skip empty line for specific banners
	if str == "shadow" || str == "thinkertoy" {
		scanner.Scan()
	}

	// Insert data into the map
	for i := 32; i < 127; i++ {
		if scanner.Scan() {
			_map[i] = InsertValue(scanner) // Ensure the map is populated correctly with each iteration
		} else {
			break
		}
	}

	defer file.Close()
}

// Apply color to a text
func applyColor(color, text string) string {
	colorCode, exists := colorCodes[color]
	if !exists {
		return text // If color doesn't exist, return original text
	}
	return colorCode + text + colorCodes["reset"]
}

// Print ASCII art and apply color
func Printing(inp, color, substring string) {
	fmt.Println("inp---->",inp)
	fmt.Println("color---->",color)
	fmt.Println("substring---->",substring)
	if inp == "\\n" {
		fmt.Println()
		return
	}

	SplArgs := strings.Split(inp, "\\n")
	//fmt.Println("-----", SplArgs)

	// Handle multiple "\n"
	if A.IsOnlyNewLine(SplArgs) {
		for i := 0; i < len(SplArgs)-1; i++ {
			fmt.Println()
		}
		return
	}

	// Process each part of the ASCII art
	for _, arg := range SplArgs {
		if arg == "" {
			fmt.Println()
			continue
		}

		// Split the argument into parts (before and after the substring)
		startIdx := strings.Index(arg, substring)
		if startIdx != -1 {
			// If substring is found, split into parts
			beforeSubstring := arg[:startIdx]
			fmt.Println("beforeSubstring-----", beforeSubstring)
			afterSubstring := arg[startIdx+len(substring):]
			fmt.Println("afterSubstring-----", afterSubstring)
			// Apply color to the substring
			coloredSubstring := applyColor(color, substring)

			// Rebuild the line with color applied only to the substring
			all :=beforeSubstring + coloredSubstring + afterSubstring
			
			fmt.Println(all)
		} else {
			// If no substring, just print the whole line without color
			fmt.Println(arg)
		}
	}
}

// Process arguments for the color flag, substring to be colored, and the main string
func processArgs(args []string) (string, string, string) {
	if len(args) < 3 || !strings.HasPrefix(args[0], "--color=") {
		fmt.Println("Usage: go run . --color=<color> <substring to be colored> [STRING]")
		os.Exit(1)
	}

	// Extract color, substring, and the full text to color
	color := strings.TrimPrefix(args[0], "--color=")
	substring := args[1]
	text := strings.Join(args[2:], " ")

	return color, substring, text
}

func main() {
	args := os.Args[1:]

	// Check and process arguments
	color, substring, text := processArgs(args)

	// Initialize map based on the banner provided
	InitMap("standard") // You can change the banner here as needed

	// Print the ASCII art with color
	Printing(text, color, substring)
}
