package main

import (
	"bufio"
	"os"

	A "asciiart-fs/utils"
	F "fmt"
	S "strings"
)

var (
	_map  = make(map[int][8]string)
	lines = [8]string{}
)

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
	// to skip first empty line
	if str == "shadow" || str == "thinkertoy" {
		scanner.Scan()
	}
	// insert data on the map
	for i := 32; i < 127; i++ {
		_map[i] = A.InsertValue(scanner)
	}

	defer file.Close()
}

func Printing(inp string) {
	if inp == "\\n" {
		F.Println()
		return
	}

	SplArgs := S.Split(inp, "\\n")

	// Fix Multiple "\n"
	if A.IsOnlyNewLine(SplArgs) {
		for i := 0; i < len(SplArgs)-1; i++ {
			F.Println()
		}
		return
	}
	// applying "\n"
	for _, arg := range SplArgs {
		if arg == "" {
			F.Println()
			continue
		}
		// Storing Data
		for _, val := range arg {
			for i := 0; i < 8; i++ {
				lines[i] += _map[int(val)][i]
			}
		}
		// Printing Data
		for i := 0; i < 8; i++ {
			F.Println(lines[i])
			lines[i] = ""
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		os.Stderr.WriteString("Err: Invalid Format | Usage: [STRING] [BANNER] !")
		return
	}

	if !A.IsValidArg(args[0]) {
		os.Stderr.WriteString("Err: Invalid Arguments !")
		return
	}

	InitMap(args[1])
	Printing(args[0])
}
