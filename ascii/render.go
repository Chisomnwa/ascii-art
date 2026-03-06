package ascii

import "os"

// func Render(text string, banner string) {
// 	bannerLines := readBanner(banner)

// 	// convert banner file to map
// 	asciiMap := buildAsciiMap(bannerLines)

// 	// print ascii art
// 	printAscii(text, asciiMap)
// }

import (
	//"fmt"
	"strings"
)

func ReadBanner(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}