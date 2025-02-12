package main

import (
	"flag"
	"fmt"

	"strings"
)

// Map of color names to their corresponding ANSI color codes (0 to 255)
var colorMap = map[string]string{
	"black": "0", "red": "1", "green": "2", "yellow": "3", "blue": "4", "magenta": "5", "cyan": "6", "white": "7",
	"lightGray": "8", "lightRed": "9", "lightGreen": "10", "lightYellow": "11", "lightBlue": "12", "lightMagenta": "13", "lightCyan": "14", "lightWhite": "15",
	"darkGray": "16", "lightGray2": "17", "darkBlue": "19", "darkGreen": "20", "lightPink": "21", "fuchsia": "22", "darkYellow": "23",
	"gray": "24", "lightGray3": "25", "lightPink2": "26", "darkBrown": "27", "lightTan": "28", "darkBeige": "29", "violet": "30", "purple": "31",
	// You can extend this map for more colors if needed...
}

// Function to print colored text using ANSI escape codes
func printColoredText(colorName string, text string) {
	// Check if the color name exists in the map
	if colorCode, exists := colorMap[colorName]; exists {
		// Apply the color using ANSI escape code
		fmt.Printf("\033[38;5;%s%m%s\033[0m\n", colorCode, text)
	} else {
		fmt.Printf("Invalid color name: %s. Please provide a valid color name.\n", colorName)
	}
}

func main() {
	// Define the flags
	colorFlag := flag.String("color", "", "The color name to apply to the text")
	flag.Parse()

	// Check if a color was provided
	if *colorFlag == "" || len(flag.Args()) < 1 {
		fmt.Println("Usage: go run . --color=<color> <substring to be colored>")
		fmt.Println("Example: go run . --color=red \"This is red text!\"")
		return
	}

	// The rest of the arguments are the text to be colored
	textToColor := strings.Join(flag.Args(), " ")

	// Print the colored text
	printColoredText(*colorFlag, textToColor)
}
