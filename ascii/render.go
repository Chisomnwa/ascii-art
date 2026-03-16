package ascii

import (
	"os"
	"fmt"
	"strings"
)

func ReadBanner(file string) ([]string, error) {
	/*
	Reads banner files
	*/
	data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		lines := strings.Split(string(data), "\n")
		return lines, err
	}

func BuildAsciiMap(lines []string) map[rune][]string {
	/*
	Converts the banner files to a map container
	*/ 
	asciiMap := make(map[rune][]string)

	char := 32

	// Store exactly eight lines of each ASCII character 
	// ignoring the blank that seperates them
	for i := 1; i < len(lines); i += 9 {
		asciiMap[rune(char)] = lines[i : i+8]
		char++
	}

	return asciiMap
}

func PrintAscii(text string, asciiMap map[rune][]string) string {
	/*
	Prints the aascii art on the command line
	*/

	// decalre what we want to return
	var result strings.Builder


	// Split the text with into new lines using "\n" as a seperator
	lines := strings.Split(text, "\\n")

	// In case "\n" appears as the first line, ignore it
	// so we don't print unnecessary leading blank line
	for i, line := range lines {

		if line == "" {
			if i != 0 {
			fmt.Println()
		}
		continue
		}

		for row := 0; row < 8; row++ {

			for _, char := range line {
				result.WriteString(asciiMap[char][row])
			}
			result.WriteString("\n")
		}
	}

	// Retrieve and return the final accumulated string
	return result.String()
}