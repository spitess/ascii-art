package main

import (
	"bufio"
	F "fmt"
	"os"
	S "strings"

	U "asciiArt/utils"
)

var (
	_map  = make(map[int][8]string)
	lines = [8]string{}
)

func InitMap() {
	file, err := os.Open("standard.txt")
	if err != nil {
		os.Stderr.WriteString("Err: opening file: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for i := 32; i < 127; i++ {
		_map[i] = U.InsertValue(scanner)
	}

	defer file.Close()
}

func Printing(arg string) {
	if arg == "\\n" {
		F.Println()
		return
	}
	args := S.Split(arg, "\\n")

	if U.Isonlynewline(args) {

		for i := 0; i < len(args)-1; i++ {
			F.Println()
		}

		return

	}

	for _, V := range args {
		if V == "" {
			F.Println()
			continue
		}

		for _, val := range V {
			for i := 0; i < 8; i++ {
				lines[i] += _map[int(val)][i]
			}
		}

		for i := 0; i < 8; i++ {
			F.Println(lines[i])
			lines[i] = ""
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Stderr.WriteString("Err: Not Enough Parameters!")
		return
	}

	if len(args) != 1 || args[0] == "" {
		return
	}
	if !U.IsValidArg(args[0]) {
		os.Stderr.WriteString("Err: Invalid Charactere!")
		return
	}
	InitMap()
	Printing(args[0])
}
