package trojansourcedetector

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// BIDIDetector returns a detector that detects bidirectional control characters.
func BIDIDetector() SingleFileDetector {
	return &bidiDetector{}
}

// bidiCharacters contains a list of bidirectional control characters from the Unicode specification
// See https://www.unicode.org/reports/tr9/tr9-42.html
var bidiCharacters = map[rune]string{
	'\u202A': "Left-to-Right Embedding",
	'\u202B': "Right-to-Left Embedding",
	'\u202D': "Left-to-Right Override",
	'\u202E': "Right-to-Left Override",
	'\u2066': "Left-to-Right Isolate",
	'\u2067': "Right-to-Left Isolate",
	'\u2068': "First Strong Isolate",
	'\u202C': "Pop Directional Formatting",
	'\u2069': "Pop Directional Isolate",

	'\u200E': "Left-to-Right Mark",
	'\u200F': "Right-to-Left Mark",
	'\u061C': "Arabic Letter Mark",
}

type bidiDetector struct {
}

func (b bidiDetector) Detect(filename string, i io.Reader) Errors {
	input := bufio.NewReader(i)
	line := uint(0)
	column := uint(0)
	e := NewErrors()
	for {
		nextRune, size, err := input.ReadRune()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				e.Add(
					ErrIOFile,
					err.Error(),
					filename,
					line+1,
					column+1,
				)
			}
			return e
		}
		if details, ok := bidiCharacters[nextRune]; ok {
			e.Add(
				ErrBIDI,
				fmt.Sprintf("%s control character", details),
				filename,
				line+1,
				column+1,
			)
		}
		if nextRune == '\n' {
			line++
			column = 0
		} else {
			column += uint(size)
		}
	}
}
